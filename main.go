package main

import (
	"log"

	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Configure CORS
	app.Use(cors.New())

	// Connect to database
	config.ConnectDB()
	config.InitJWTConfig()

	// Setup routes
	api := app.Group("/api/v1")
	routes.SetupAuthRoutes(api)
	routes.SetupSubscriptionRoutes(api)

	log.Fatal(app.Listen(":3000"))
}
