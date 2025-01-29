package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go-restful-api/utils"
)

// AuthMiddleware checks the Authorization header for a valid token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if the token starts with "Bearer "
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader { // If Bearer is missing
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Add user claims to the context for later use
		c.Set("user", claims)

		// Continue to the next handler
		c.Next()
	}
}
