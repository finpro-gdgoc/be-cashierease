package handlers

import (
	"cashierease/internal/models"
	"cashierease/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCoupon(c *gin.Context) {
	var coupon models.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.CreateCoupon(&coupon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coupon"})
		return
	}

	c.JSON(http.StatusCreated, coupon)
}

func GetAllCoupons(c *gin.Context) {
	coupons, err := repositories.GetAllCoupons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get coupons"})
		return
	}
	c.JSON(http.StatusOK, coupons)
}

func GetCouponByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	coupon, err := repositories.GetCouponByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}
	c.JSON(http.StatusOK, coupon)
}

func UpdateCoupon(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingCoupon, err := repositories.GetCouponByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
		return
	}

	if err := c.ShouldBindJSON(&existingCoupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateCoupon(&existingCoupon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update coupon"})
		return
	}

	c.JSON(http.StatusOK, existingCoupon)
}

func DeleteCoupon(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := repositories.DeleteCoupon(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete coupon"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Coupon deleted successfully"})
}