package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/controllers"
	"github.com/chandra-devs/subscription_app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ReadBufferSize:  16 * 1024,
		WriteBufferSize: 16 * 1024,
		BodyLimit:       10 * 1024 * 1024,
	})

	// Configure CORS
	app.Use(cors.New())

	// Connect to database
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := config.CloseDB(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Initialize JWT configuration
	config.InitJWTConfig()

	// Setup routes
	routes.SetupRoutes(app)

	// Create channel for graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-shutdown
	log.Println("Shutting down server...")

	// Cleanup and graceful shutdown
	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
}

// routes/setup.go - update only the SetupUserRoutes function
func SetupUserRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser) // Add the missing DELETE route
	users.Get("/:id/subscriptions", controllers.GetUserSubscriptions)
	users.Post("/:id/subscriptions", controllers.AddSubscription)
}
