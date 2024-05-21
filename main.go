package main

import (
	"go-auth/database"
	"go-auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to the database
	database.Connect()

	// Create a new Fiber instance
	app := fiber.New()

	// Setup CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	// Setup your routes after setting up CORS
	routes.Setup(app)

	// Start the server
	app.Listen(":8000")
}
