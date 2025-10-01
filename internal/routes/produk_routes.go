package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProdukRoutes(router *gin.RouterGroup) {
	produkRoutes := router.Group("/produk")
	{
		produkRoutes.GET("/", handlers.GetAllProduk)
		produkRoutes.GET("/:id", handlers.GetProdukById)

		produkRoutes.POST("/", middleware.AuthMiddleware(), handlers.CreateProduk)
		produkRoutes.PATCH("/:id", middleware.AuthMiddleware(), handlers.UpdateProduk)
		produkRoutes.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteProduk)
	}
}