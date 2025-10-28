package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTokoRoutes(router *gin.RouterGroup) {
	tokoRoutes := router.Group("/toko")
	{
		tokoRoutes.GET("/", handlers.GetToko)
		
		protected := tokoRoutes.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.PATCH("/", handlers.UpdateToko)
		}
	}
}