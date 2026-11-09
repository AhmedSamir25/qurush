package database

import (
	"fmt"
	"log"
	"os"
	"qurush/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func ConnectionDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatalf("Database Connection Failed: %v", err)
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("migration failed:", err)
	}
	DBConnection = db
}
