package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go-restful-api/config"
	"go-restful-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProfileByUserID godoc
// @Summary Create a new profile
// @Description Add a new profile to the database. User ID is extracted from the token.
// @Tags profiles
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.Profile true "Profile details"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profiles [post]
func CreateProfileByUserID(c *gin.Context) {
	profileCollection := config.GetCollection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var profile models.Profile

	// Ambil UserID dari Token
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi ke struct Claims
	claims, ok := userData.(*models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
		return
	}

	userIDStr := claims.UserID

	// Validasi User ID
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
		return
	}

	// Cek apakah profil sudah ada untuk user ini
	var existingProfile models.Profile
	err = profileCollection.FindOne(ctx, bson.M{"user_id": userObjectID}).Decode(&existingProfile)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already has a profile"})
		return
	}

	// Bind JSON input ke struct profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Isi data profil baru
	profile.ID = primitive.NewObjectID()
	profile.UserID = userObjectID
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	// Simpan ke database
	_, err = profileCollection.InsertOne(ctx, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully", "profile": profile})
}

// GetProfileByUserID godoc
// @Summary Get profile of authenticated user
// @Description Retrieve profile details using the User ID from token.
// @Tags profiles
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.Profile
// @Failure 404 {object} map[string]string
// @Router /profiles [get]
func GetProfileByUserID(c *gin.Context) {
	profileCollection := config.GetCollection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ambil UserID dari Token
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi ke struct Claims
	claims, ok := userData.(*models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
		return
	}

	userIDStr := claims.UserID

	// Validasi User ID
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
		return
	}

	// Cari profil berdasarkan user_id
	var profile models.Profile
	err = profileCollection.FindOne(ctx, bson.M{"user_id": userObjectID}).Decode(&profile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// UpdateProfileByUserID godoc
// @Summary Update profile of authenticated user
// @Description Update profile fields of the logged-in user
// @Tags profiles
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body models.Profile true "Updated Profile details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profiles [put]
func UpdateProfileByUserID(c *gin.Context) {
	profileCollection := config.GetCollection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ambil UserID dari Token
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi ke struct Claims
	claims, ok := userData.(*models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
		return
	}

	userIDStr := claims.UserID

	// Validasi User ID
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
		return
	}

	// Cek apakah profil ada
	var existingProfile models.Profile
	err = profileCollection.FindOne(ctx, bson.M{"user_id": userObjectID}).Decode(&existingProfile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Bind JSON input ke struct sementara
	var updatedProfile models.Profile
	if err := c.ShouldBindJSON(&updatedProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buat update object hanya untuk field yang diisi
	updateFields := bson.M{
		"updated_at": time.Now(),
	}

	if updatedProfile.Bio != "" {
		updateFields["bio"] = updatedProfile.Bio
	}
	if updatedProfile.Avatar != "" {
		updateFields["avatar"] = updatedProfile.Avatar
	}

	// Lakukan update di database
	_, err = profileCollection.UpdateOne(ctx, bson.M{"user_id": userObjectID}, bson.M{"$set": updateFields})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// DeleteProfileByUserID godoc
// @Summary Delete profile of authenticated user
// @Description Remove profile of the logged-in user
// @Tags profiles
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /profiles [delete]
func DeleteProfileByUserID(c *gin.Context) {
	profileCollection := config.GetCollection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ambil UserID dari Token
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Konversi ke struct Claims
	claims, ok := userData.(*models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
		return
	}

	userIDStr := claims.UserID

	// Validasi User ID
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
		return
	}

	// Cek apakah profil ada
	var existingProfile models.Profile
	err = profileCollection.FindOne(ctx, bson.M{"user_id": userObjectID}).Decode(&existingProfile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Hapus profil dari database
	_, err = profileCollection.DeleteOne(ctx, bson.M{"user_id": userObjectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}