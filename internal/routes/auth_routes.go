package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"
	"cashierease/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)

		adminOnly := authRoutes.Group("/")
		adminOnly.Use(middleware.AuthMiddleware())
		adminOnly.Use(middleware.RoleMiddleware(models.AdminRole))
		{
			adminOnly.GET("/", handlers.GetAllUsers)
			adminOnly.GET("/:id", handlers.GetUserByID)
			adminOnly.PATCH("/:id", handlers.UpdateUser)
			adminOnly.DELETE("/:id", handlers.DeleteUser)
		}
	}
}