package routes

import (
	"github.com/gofiber/fiber/v2"
	"user/internal/handlers"
	"user/internal/middleware"
)

func SetupV1Routes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/login", handlers.Login)

	api.Get("/users", middleware.AccessMiddleware("admin"), handlers.Index)
	api.Post("/users", middleware.AccessMiddleware("admin"), handlers.Store)
	api.Get("/users/:id", middleware.AccessMiddleware("admin"), handlers.Detail)
}
