package config

import (
	"fmt"
	"log"
	"os"
	"cashierease/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection successful.")

	database.AutoMigrate(&models.Produk{}) 
	fmt.Println("Database migration successful.")

	database.AutoMigrate(&models.Produk{}, &models.User{}) 
	fmt.Println("Database migration successful.")

	DB = database
}