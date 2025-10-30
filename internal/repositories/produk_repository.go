package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"
)

func CreateProduk(produk *models.Produk) error {
	result := config.DB.Create(produk)
	return result.Error
}

func GetAllProduk() ([]models.Produk, error) {
	var produks []models.Produk
	result := config.DB.Find(&produks)
	return produks, result.Error
}

func GetProdukById(id uint) (models.Produk, error) {
	var produk models.Produk
	result := config.DB.First(&produk, id)
	return produk, result.Error
}

func UpdateProduk(produk *models.Produk) error {
	result := config.DB.Save(produk)
	return result.Error
}

func DeleteProduk(id uint) error {
	result := config.DB.Delete(&models.Produk{}, id)
	return result.Error
}