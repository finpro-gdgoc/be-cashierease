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

func SearchProdukByName(name string) ([]models.Produk, error) {
	var produks []models.Produk
	query := "%" + name + "%"
	err := config.DB.Where("nama_produk ILIKE ?", query).Find(&produks).Error
	return produks, err
}

func GetProdukBySlug(slug string) (models.Produk, error) {
	var produk models.Produk
	err := config.DB.Where("slug_produk = ?", slug).First(&produk).Error
	return produk, err
}