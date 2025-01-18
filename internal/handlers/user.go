package handlers

import (
	"github.com/gofiber/fiber/v2"
	"user/internal/database"
	"user/internal/models"
)

func UserIndex(c *fiber.Ctx) error {
	users := []models.User{}
	database.DBConn.Find(&users)

	res := c.JSON(users)
	return res
}

func UserDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	database.DBConn.First(&user, id)

	return c.JSON(user)
}
