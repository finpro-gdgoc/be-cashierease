package handlers

import (
	"cashierease/internal/models"
	"cashierease/internal/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	OrderItems []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"order_items"`
	Coupon struct {
		CouponID uint `json:"couponId"`
	} `json:"coupon"`
	PaymentMethod string `json:"payment_method"`
}

func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var totalPrice float64 = 0
	var orderItemsDetail []models.OrderItem
	const tax float64 = 0.1

	for _, item := range input.OrderItems {
		produk, err := repositories.GetProdukById(item.ProductID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk not found: " + produk.NamaProduk})
			return
		}

		totalPrice += produk.HargaProduk * float64(item.Quantity)

		orderItemsDetail = append(orderItemsDetail, models.OrderItem{
			ProductID:   item.ProductID,
			ProductName: produk.NamaProduk,
			Quantity:    item.Quantity,
		})
	}

	var couponDetails models.CouponDetails
	var totalPriceDiscount float64 = 0
	var discountAmount float64 = 0

	if input.Coupon.CouponID != 0 {
		coupon, err := repositories.GetCouponByID(input.Coupon.CouponID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
			return
		}

		couponDetails = models.CouponDetails{
			CouponID:      coupon.ID,
			KodeCoupon:    coupon.KodeCoupon,
			BesarDiscount: coupon.BesarDiscount,
		}
		discountAmount = coupon.BesarDiscount
		totalPriceDiscount = totalPrice - (totalPrice * discountAmount)
	}

	var finalPriceWithTax float64
	if totalPriceDiscount > 0 {
		finalPriceWithTax = totalPriceDiscount + (totalPriceDiscount * tax)
	} else {
		finalPriceWithTax = totalPrice + (totalPrice * tax)
	}

	order := models.Order{
		OrderDate:            time.Now(),
		TotalPrice:           totalPrice,
		TotalPriceWithDiscount: totalPriceDiscount,
		TotalPriceWithTax:    finalPriceWithTax,
		PaymentMethod:        input.PaymentMethod,
		Tax:                  tax,
		DiscountAmount:       discountAmount,
		Coupon:               couponDetails,
		OrderItems:           orderItemsDetail,
	}

	if err := repositories.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetAllOrders(c *gin.Context) {
	orders, err := repositories.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}