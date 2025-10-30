package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"
	"time"
)

func CreateCoupon(coupon *models.Coupon) error {
	return config.DB.Create(coupon).Error
}

func GetAllCoupons() ([]models.Coupon, error) {
	var coupons []models.Coupon
	err := config.DB.Find(&coupons).Error
	return coupons, err
}

func GetCouponByID(id uint) (models.Coupon, error) {
	var coupon models.Coupon
	err := config.DB.First(&coupon, id).Error
	return coupon, err
}

func UpdateCoupon(coupon *models.Coupon) error {
	return config.DB.Save(coupon).Error
}

func DeleteCoupon(id uint) error {
	return config.DB.Delete(&models.Coupon{}, id).Error
}

func GetCouponByCode(kode string) (models.Coupon, error) {
	var coupon models.Coupon
	err := config.DB.Where("kode_coupon = ?", kode).First(&coupon).Error
	return coupon, err
}

func GetActiveCoupons() ([]models.Coupon, error) {
	var coupons []models.Coupon
	err := config.DB.Where("akhir_coupon > ?", time.Now()).Find(&coupons).Error
	return coupons, err
}