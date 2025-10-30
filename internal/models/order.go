package models

import (
	"time"

	"gorm.io/gorm"
)

type CouponDetails struct {
	CouponID      uint    `json:"coupon_id"`
	KodeCoupon    string  `json:"kode_coupon"`
	BesarDiscount float64 `json:"besar_discount"`
}

type Order struct {
	gorm.Model
	OrderDate            time.Time     `json:"order_date"`
	TotalPrice           float64       `gorm:"not null" json:"total_price"`
	TotalPriceWithDiscount float64     `json:"total_price_with_discount"`
	TotalPriceWithTax    float64       `gorm:"not null" json:"total_price_with_tax"`
	PaymentMethod        string        `gorm:"not null" json:"payment_method"`
	Tax                  float64       `gorm:"not null;default:0.1" json:"tax"`
	DiscountAmount       float64       `json:"discount_amount"`
	Coupon               CouponDetails `gorm:"embedded" json:"coupon"`
	OrderItems           []OrderItem   `gorm:"foreignKey:OrderID" json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	OrderID     uint   `json:"-"`
	ProductID   uint   `gorm:"not null" json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `gorm:"not null" json:"quantity"`
}