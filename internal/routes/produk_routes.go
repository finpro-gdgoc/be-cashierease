package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"
	"cashierease/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupProdukRoutes(router *gin.RouterGroup) {
	produkRoutes := router.Group("/produk")
	{
		produkRoutes.GET("/", handlers.GetAllProduk)
		produkRoutes.GET("/:id", handlers.GetProdukById)

		produkRoutes.GET("/search", handlers.SearchProdukByName)
		produkRoutes.GET("/slug/:slug", handlers.GetProdukBySlug)

		adminOnly := produkRoutes.Group("/")
		adminOnly.Use(middleware.AuthMiddleware()) 
		adminOnly.Use(middleware.RoleMiddleware(models.AdminRole)) 
		{
			adminOnly.POST("/", handlers.CreateProduk)
			adminOnly.PATCH("/:id", handlers.UpdateProduk)
			adminOnly.DELETE("/:id", handlers.DeleteProduk)
			
			adminOnly.PATCH("/:id/upload", handlers.UploadGambarProduk)
		}
	}
}