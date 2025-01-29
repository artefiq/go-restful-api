package main

import (
	"github.com/gin-gonic/gin"
	_ "go-restful-api/docs" // Import generated Swagger docs
	"go-restful-api/config"
	"go-restful-api/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go RESTful API Example
// @version 1.0
// @description A simple RESTful API with MongoDB and Swagger documentation.
// @termsOfService http://example.com/terms/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Connect to MongoDB
	config.ConnectDatabase()

	// Set up Gin router
	router := gin.Default()

	// Add Swagger UI at /api/v1/swagger/*
	swaggerURL := ginSwagger.URL("http://localhost:8080/api/v1/swagger/doc.json") // Adjust Swagger base path
	router.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// Group routes under /api/v1
	api := router.Group("/api/v1")
	{
		// Register user routes within the /api/v1 group
		routes.RegisterUserRoutes(api)
		routes.RegiterProfileRoutes(api)
	}

	// Start the server
	router.Run(":8080")
}
