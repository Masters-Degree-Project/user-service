package middleware

import (
	"encoding/base64"
	"fmt"
	"user/pkg/config"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
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
