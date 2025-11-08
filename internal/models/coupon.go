package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Coupon struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	KodeCoupon    string    `gorm:"unique;not null" json:"kode_coupon"`
	AwalCoupon    time.Time `gorm:"not null" json:"awal_coupon"`
	AkhirCoupon   time.Time `gorm:"not null" json:"akhir_coupon"`
	BesarDiscount float64   `gorm:"not null" json:"besar_discount"`
	Deskripsi     string    `json:"deskripsi"`
	PaymentMethod string    `gorm:"not null" json:"payment_method"`
}

func (coupon *Coupon) BeforeCreate(tx *gorm.DB) (err error) {
	coupon.ID = uuid.New()
	return
}