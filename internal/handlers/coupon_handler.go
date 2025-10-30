package handlers

import (
	"cashierease/internal/models"
	"cashierease/internal/repositories"
	"net/http"
	"strconv"
	"time"

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

type ValidateCouponInput struct {
	KodeCoupon    string `json:"kode_coupon" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

func ValidateCoupon(c *gin.Context) {
	var input ValidateCouponInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupon, err := repositories.GetCouponByCode(input.KodeCoupon)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kode kupon tidak ditemukan"})
		return
	}

	if time.Now().After(coupon.AkhirCoupon) {
		c.JSON(http.StatusGone, gin.H{"error": "Kupon telah kedaluwarsa"})
		return
	}

	if time.Now().Before(coupon.AwalCoupon) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kupon belum berlaku"})
		return
	}

	if coupon.PaymentMethod != input.PaymentMethod {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kupon tidak berlaku untuk metode pembayaran ini"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kupon valid",
		"data":    coupon,
	})
}

func GetActiveCoupons(c *gin.Context) {
	coupons, err := repositories.GetActiveCoupons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get active coupons"})
		return
	}

	c.JSON(http.StatusOK, coupons)
}