package middleware

import (
	"encoding/base64"
	"fmt"
	"user/internal/database"
	"user/internal/models"
	"user/pkg/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(successHandler *fiber.Handler) fiber.Handler {
	var handler = func(ctx *fiber.Ctx) error {
		return ctx.Next()
	}
	if successHandler != nil {
		handler = *successHandler
	}

	secret := []byte(config.Config("JWT_SECRET"))
	encodedKey := base64.StdEncoding.EncodeToString(secret)

	decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		fmt.Println("Secret key decoding error:", err)
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: "HS256",
			Key:    decodedKey,
		},
		ErrorHandler: jwtError,
		SuccessHandler: func(ctx *fiber.Ctx) error {
			token := ctx.Locals("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)

			userId, ok := claims["id"].(float64)
			if !ok {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Invalid token",
				})
			}

			var user models.User
			result := database.DBConn.First(&user, userId)
			if result.Error != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Unauthenticated",
				})
			}

			ctx.Locals("user", &user)

			if handler != nil {
				return handler(ctx)
			}

			return ctx.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}

	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT",
		"data":    err.Error(),
	})
}
