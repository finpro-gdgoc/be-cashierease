package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	CashierRole UserRole = "cashier"
)

type User struct {
	gorm.Model
	Nama         string   `gorm:"not null"`
	NomorPegawai string   `gorm:"unique;not null"`
	Password     string   `gorm:"not null"`
	Role         UserRole `gorm:"type:varchar(10);not null;default:'user'"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}