package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware() fiber.Handler {
	return JwtMiddleware(nil)
}
