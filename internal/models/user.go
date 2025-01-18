package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model

	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Email       string `gorm:"index" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	Role        string `json:"role"`
	Picture     string `json:"picture"`
	CreatedById string `json:"created_by"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}
