package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.RouterGroup) {
	orderRoutes := router.Group("/order")
	{
		orderRoutes.GET("/", handlers.GetAllOrders)
		
		protected := orderRoutes.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/", handlers.CreateOrder)
		}

		statisticsRoutes := orderRoutes.Group("/statistics")
		{
			statisticsRoutes.GET("/", handlers.GetAllStatistics)
			statisticsRoutes.GET("/pendapatan", handlers.GetAllPendapatan)
			statisticsRoutes.GET("/pelanggan", handlers.GetAllPelanggan)
			statisticsRoutes.GET("/popularMenu/:mon", handlers.GetPopularMenu)
		}
	}
}