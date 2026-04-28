package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // pointer trỏ tới database instance, public

func Connect() { // khởi tạo kết nối tới database

	// load .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	// đọc biến môi trường từ os
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")
	port := os.Getenv("DB_PORT")

	// config của database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Mở kết nối tới database
	// Tham số thứ nhất là driver và cấu hình kết nối
	// Tham số thứ hai là config của GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db // nếu mở thành công, gán connection đó vào biến global
	fmt.Println("Database connected")
}
