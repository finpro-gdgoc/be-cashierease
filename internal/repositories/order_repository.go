package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"
)

func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Preload("OrderItems").Find(&orders).Error
	return orders, err
}