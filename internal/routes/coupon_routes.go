package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupCouponRoutes(router *gin.RouterGroup) {
	couponRoutes := router.Group("/coupon")
	{
		couponRoutes.GET("/", handlers.GetAllCoupons)
		couponRoutes.GET("/:id", handlers.GetCouponByID)

		// Rute yang memerlukan autentikasi
		protected := couponRoutes.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/", handlers.CreateCoupon)
			protected.PATCH("/:id", handlers.UpdateCoupon)
			protected.DELETE("/:id", handlers.DeleteCoupon)
		}
	}
}