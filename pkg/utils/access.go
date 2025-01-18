package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func HasAccess(c *fiber.Ctx, role string) bool {
	userRole := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["role"].(string)
	return userRole == role
}
