package routes

import (
	"github.com/gin-gonic/gin"
	"go-restful-api/controllers"
	"go-restful-api/middleware"
)

// RegisterUserRoutes registers routes for user-related operations
func RegisterUserRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	{
		// Public route: Create user (registration)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.POST("/login", controllers.LoginUser)

		// Protected routes: Require authentication
		userRoutes.Use(middleware.AuthMiddleware()) // Apply AuthMiddleware to all routes below

		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
