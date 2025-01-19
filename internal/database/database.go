package database

import (
	"fmt"
	"log"
	"os"
	"user/internal/models"
	"user/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDb() {
	dbUser := config.Config("DB_USER")
	dbPassword := config.Config("DB_PASSWORD") // Can be empty
	dbHost := config.Config("DB_HOST")
	dbPort := config.Config("DB_PORT")
	dbName := config.Config("DB_NAME")

	// Validate only essential environment variables
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("Required database configuration environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) are not set properly")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		dbUser, dbPassword, dbHost, dbPort, dbName)

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
