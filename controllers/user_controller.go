package controllers

import (
	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if result := config.DB.Preload("Subscriptions").Find(&users); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if result := config.DB.Preload("Subscriptions").First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if result := config.DB.Create(&user); result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	config.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	config.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

func AddSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if result := config.DB.First(&user, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	subscription := new(models.Subscription)
	if err := c.BodyParser(subscription); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	subscription.UserID = user.ID
	if result := config.DB.Create(&subscription); result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to create subscription"})
	}

	return c.Status(201).JSON(subscription)
}

func GetUserSubscriptions(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscriptions []models.Subscription
	if result := config.DB.Where("user_id = ?", id).Find(&subscriptions); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Subscriptions not found"})
	}
	return c.JSON(subscriptions)
}
