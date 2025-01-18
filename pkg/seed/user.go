package seed

import (
	"user/internal/database"
	"user/internal/models"
)

func AdminUserIfDoesntExist() {
	var adminUserEmail = "admin@example.com"
	var existingUser models.User

	// Check if admin user exists
	if err := database.DBConn.Where("email = ?", adminUserEmail).First(&existingUser).Error; err != nil {
		// Create admin user only if it doesn't exist
		pass, _ := models.HashPassword("Test123123")

		adminUser := models.User{
			Name:     "John",
			Lastname: "Doe",
			Email:    adminUserEmail,
			Password: pass,
			Role:     "admin",
		}

		database.DBConn.Create(&adminUser)
	}
}
