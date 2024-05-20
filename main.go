package main

import (
	"go-auth/database"
	"go-auth/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // or specify specific origins like "http://example.com"
		AllowCredentials: false, // set this to false when using a wildcard for origins
	}))
	

	routes.Setup(app)

	app.Listen(":8000")
}