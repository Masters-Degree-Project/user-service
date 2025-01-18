package repositories

import (
	"errors"
	"gorm.io/gorm"
	"user/internal/database"
	"user/internal/models"
)

func GetUserById(id int) (*models.User, error) {
	db := database.DBConn
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(e string) (*models.User, error) {
	db := database.DBConn
	var user models.User

	if err := db.Where(&models.User{Email: e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
