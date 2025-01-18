package middleware

import (
	"user/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func AccessMiddleware(role string) fiber.Handler {
	successHandler := func(ctx *fiber.Ctx) error {
		hasRoleAccess := utils.HasAccess(ctx, role)
		if !hasRoleAccess {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "access-error",
				"message": "Access denied! Required role: " + role,
			})
		}

		return ctx.Next()
	}

	return JwtMiddleware(&successHandler)
}
