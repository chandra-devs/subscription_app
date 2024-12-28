// main.go
package main

import (
	"log"

	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/controllers"
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

	// Setup all routes using the existing SetupRoutes function
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
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
