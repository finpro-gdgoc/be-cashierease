package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CouponDetails struct {
	CouponID      uuid.UUID `json:"coupon_id"`
	KodeCoupon    string  `json:"kode_coupon"`
	BesarDiscount float64 `json:"besar_discount"`
}

type Order struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

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
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	OrderID     uuid.UUID `json:"-"`
	ProductID   uuid.UUID `gorm:"not null" json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `gorm:"not null" json:"quantity"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID = uuid.New()
	return
}

func (orderItem *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	orderItem.ID = uuid.New()
	return
}