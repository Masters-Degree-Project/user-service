package main

import (
	"fmt"
	"log"
	"user/internal/database"
	"user/internal/routes"
	"user/pkg/config"
	"user/pkg/consul"
	"user/pkg/seed"

	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	myFigure := figure.NewColorFigure("User Service", "", "green", true)
	myFigure.Print()

	log.Println("Starting User Service")

	log.Println("Consul registration...")
	if err := consul.RegisterService(); err != nil {
		log.Fatalf("Consul registration failed: %v", err)
	}

	database.ConnectDb()

	log.Println("Admin user checking...")
	seed.AdminUserIfDoesntExist()

	log.Println("Starting Fiber App")

	// Initialize a new Fiber app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	routes.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Start the server on port
	port := config.Config("SERVICE_PORT")
	go log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
