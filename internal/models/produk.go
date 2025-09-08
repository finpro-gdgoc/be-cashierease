package models

import (
	"time"
	"gorm.io/gorm"
)

type TipeProduk string

const (
	Makanan TipeProduk = "makanan"
	Minuman TipeProduk = "minuman"
	Snack   TipeProduk = "snack"
)

type Produk struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	NamaProduk  string    `gorm:"not null" json:"nama_produk"`
	HargaProduk float64   `gorm:"not null" json:"harga_produk"`
	StokProduk  int       `gorm:"not null" json:"stok_produk"`
	TipeProduk  TipeProduk `gorm:"type:varchar(20);not null" json:"tipe_produk"`
	SlugProduk  string    `gorm:"unique" json:"slug_produk"`
	GambarProduk string   `json:"gambar_produk"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}