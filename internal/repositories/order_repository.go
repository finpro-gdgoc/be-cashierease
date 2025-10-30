package repositories

import (
	"cashierease/config"
	"cashierease/internal/models"
	"time"
)

func CreateOrder(order *models.Order) error {
	return config.DB.Create(order).Error
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func GetOrdersByDateRange(start time.Time, end time.Time) ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Preload("OrderItems").
		Where("order_date BETWEEN ? AND ?", start, end).
		Find(&orders).Error
	return orders, err
}