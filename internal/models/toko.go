package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Toko struct {
	ID       uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	NamaToko string `gorm:"default:'Kasir Kilat'" json:"nama_toko"`
}

func (toko *Toko) BeforeCreate(tx *gorm.DB) (err error) {
	toko.ID = uuid.New()
	return
}