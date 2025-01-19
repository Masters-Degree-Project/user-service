package handlers

import (
	"github.com/gofiber/fiber/v2"
	"user/internal/database"
	"user/internal/models"
)

func Index(c *fiber.Ctx) error {
	var users []models.User
	database.DBConn.Find(&users)

	res := c.JSON(users)
	return res
}

func Store(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

func Detail(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	database.DBConn.First(&user, id)

	return c.JSON(user)
}
