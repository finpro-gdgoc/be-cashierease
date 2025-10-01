package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindUserByNomorPegawai(nomorPegawai string) (models.User, error) {
	var user models.User
	err := config.DB.Where("nomor_pegawai = ?", nomorPegawai).First(&user).Error
	return user, err
}