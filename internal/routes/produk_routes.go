package routes

import (
	"cashierease/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupProdukRoutes(router *gin.RouterGroup) {
	produkRoutes := router.Group("/produk")
	{
		produkRoutes.POST("/", handlers.CreateProduk)
		produkRoutes.GET("/", handlers.GetAllProduk)
		produkRoutes.GET("/:id", handlers.GetProdukById)
		produkRoutes.PATCH("/:id", handlers.UpdateProduk)
		produkRoutes.DELETE("/:id", handlers.DeleteProduk)
	}
}