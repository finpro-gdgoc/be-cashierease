package routes

import (
	"cashierease/internal/handlers"
	"cashierease/internal/middleware"
	"cashierease/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupCouponRoutes(router *gin.RouterGroup) {
	couponRoutes := router.Group("/coupon")
	{
		couponRoutes.GET("/", handlers.GetAllCoupons)
		couponRoutes.GET("/:id", handlers.GetCouponByID)

		adminOnly := couponRoutes.Group("/")
		adminOnly.Use(middleware.AuthMiddleware())
		adminOnly.Use(middleware.RoleMiddleware(models.AdminRole))
		{
			adminOnly.POST("/", handlers.CreateCoupon)
			adminOnly.PATCH("/:id", handlers.UpdateCoupon)
			adminOnly.DELETE("/:id", handlers.DeleteCoupon)
		}
	}
}