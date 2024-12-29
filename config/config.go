package config

import (
	"fmt"
	"log"
	"os"

	"github.com/cholid97/go-kredit/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envUserDB := os.Getenv("DB_USER")
	envPassDB := os.Getenv("DB_PASSWORD")
	envHostDB := os.Getenv("DB_HOST")
	envPortDB := os.Getenv("DB_PORT")
	envDBName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		envUserDB, envPassDB, envHostDB, envPortDB, envDBName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{}, &models.Contract{}, &models.Limit{})

	DB = database
	fmt.Println("Database connection established successfully!")
}
