package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AccessMiddleware(role string) fiber.Handler {
	successHandler := func(ctx *fiber.Ctx) error {
		//hasUsersViewAccess := Utils.HasAccess(ctx, role)
		//if !hasUsersViewAccess {
		//	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		//		"status":  "access-error",
		//		"message": "You do not have access to this resource! You need role (" + role + ") to access.",
		//	})
		//}

		return ctx.Next()
	}

	return AuthMiddleware(successHandler)
}
