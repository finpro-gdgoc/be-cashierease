package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"

	"github.com/google/uuid"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindUserByNomorPegawai(nomorPegawai string) (models.User, error) {
	var user models.User
	err := config.DB.Where("nomor_pegawai = ?", nomorPegawai).First(&user).Error
	return user, err
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func GetUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return user, err
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uuid.UUID) error {
	return config.DB.Delete(&models.User{}, id).Error
}