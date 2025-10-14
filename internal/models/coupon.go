package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	KodeCoupon    string    `gorm:"unique;not null" json:"kode_coupon"`
	AwalCoupon    time.Time `gorm:"not null" json:"awal_coupon"`
	AkhirCoupon   time.Time `gorm:"not null" json:"akhir_coupon"`
	BesarDiscount float64   `gorm:"not null" json:"besar_discount"`
	Deskripsi     string    `json:"deskripsi"`
	PaymentMethod string    `gorm:"not null" json:"payment_method"`
}