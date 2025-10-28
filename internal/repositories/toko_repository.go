package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"
)

func GetToko() (models.Toko, error) {
	var toko models.Toko
	
	if err := config.DB.FirstOrCreate(&toko).Error; err != nil {
		return toko, err
	}
	return toko, nil
}

func UpdateToko(toko *models.Toko) error {
	return config.DB.Model(&toko).Updates(toko).Error
}