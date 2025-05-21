package repository

import (
	"coupon-system/config"
	"coupon-system/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	c := config.AppConfig
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	err = DB.AutoMigrate(
		&models.Coupon{})

	if err != nil {
		log.Fatal("Error automigrating tables: ", err)
	}
	log.Println("Database connection successfully established")
}
