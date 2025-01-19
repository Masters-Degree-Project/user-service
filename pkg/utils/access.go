package utils

import (
	"user/internal/models"

	"github.com/gofiber/fiber/v2"
)

func HasAccess(c *fiber.Ctx, role string) bool {
	user := c.Locals("user").(*models.User)
	return user.Role == role
}
