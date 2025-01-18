package database

import (
	"log"
	"os"
	"user/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {
	dsn := "root:@tcp(127.0.0.1:3306)/user_service?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate User model. \n", err)
		os.Exit(2)
	}

	DBConn = db
}
