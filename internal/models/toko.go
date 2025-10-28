package models

import "gorm.io/gorm"

type Toko struct {
	gorm.Model
	NamaToko string `gorm:"default:'Kasir Kilat'" json:"nama_toko"`
}