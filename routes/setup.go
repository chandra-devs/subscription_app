package routes

import (
	"github.com/chandra-devs/subscription_app/controllers"
	"github.com/chandra-devs/subscription_app/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all application routes
func SetupRoutes(app *fiber.App) {
	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Serve static files from the "docs" directory
	app.Static("/docs", "./docs")

	// Welcome screen
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("public/welcome.html")
	})

	// Ping route
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	// Generate PDF route
	app.Get("/generate-pdf", controllers.GeneratePDF)

	api := app.Group("/api/v1")

	// Setup all route groups
	SetupAuthRoutes(api)
	SetupUserRoutes(api)
	SetupSubscriptionRoutes(api)
	SetupPlanRoutes(api)
}

// SetupAuthRoutes configures authentication routes
func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}

// SetupUserRoutes configures user management routes
func SetupUserRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}

// SetupSubscriptionRoutes configures subscription management routes
func SetupSubscriptionRoutes(router fiber.Router) {
	subscriptions := router.Group("/subscriptions")
	subscriptions.Get("/user/:userId", handlers.GetUserSubscriptions)
	subscriptions.Post("/subscribe", handlers.SubscribeUser)
	subscriptions.Post("/", handlers.CreateSubscription)
}

// SetupPlanRoutes configures plan management routes
func SetupPlanRoutes(router fiber.Router) {
	plans := router.Group("/plans")
	plans.Get("/", handlers.GetPlans)
	plans.Get("/:id", handlers.GetPlanByID)
	plans.Post("/", handlers.CreatePlan)
}
