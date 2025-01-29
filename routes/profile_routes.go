package routes

import (
	"github.com/gin-gonic/gin"
	"go-restful-api/controllers"
	"go-restful-api/middleware"
)

func RegiterProfileRoutes(api *gin.RouterGroup) {
	profileRoutes := api.Group("/profiles")
	{
		// Protected route: Require Authenticated
		profileRoutes.Use(middleware.AuthMiddleware())

		profileRoutes.POST("/",controllers.CreateProfileByUserID)
		profileRoutes.GET("/", controllers.GetProfileByUserID)
		profileRoutes.PUT("/", controllers.UpdateProfileByUserID)
		profileRoutes.DELETE("/", controllers.DeleteProfileByUserID)
	}
}