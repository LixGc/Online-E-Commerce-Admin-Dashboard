package config

import (
	"e-commerce/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=E-Commerce-Golang password=Postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.TransactionHistory{})
	return db
}
