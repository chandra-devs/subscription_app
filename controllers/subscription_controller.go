package controllers

import (
	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/models"
	"github.com/gofiber/fiber/v2"
)

func CreateSubscription(c *fiber.Ctx) error {
	subscription := new(models.Subscription)
	if err := c.BodyParser(subscription); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
		})
	}

	if result := config.DB.Create(&subscription); result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create subscription",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Subscription created successfully",
		"data":    subscription,
	})
}

func GetSubscriptions(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 50)
	if limit > 100 {
		limit = 100 // Prevent memory exhaustion from large queries
	}

	offset := (page - 1) * limit
	var total int64
	var subscriptions []models.Subscription

	// Get total count with a separate query
	if err := config.DB.Model(&models.Subscription{}).Count(&total).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch subscriptions count",
		})
	}

	// Get paginated data
	if result := config.DB.
		Preload("User"). // Use preload instead of eager loading
		Preload("Plan"). // Preload related data
		Limit(limit).
		Offset(offset).
		Find(&subscriptions); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to fetch subscriptions",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"subscriptions": subscriptions,
			"pagination": fiber.Map{
				"current_page": page,
				"total_pages":  (total + int64(limit) - 1) / int64(limit),
				"total_items":  total,
				"per_page":     limit,
			},
		},
	})
}

func GetSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscription models.Subscription
	if result := config.DB.First(&subscription, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Subscription not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   subscription,
	})
}

func UpdateSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	subscription := new(models.Subscription)

	if result := config.DB.First(&subscription, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Subscription not found",
		})
	}

	if err := c.BodyParser(subscription); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
		})
	}

	config.DB.Save(&subscription)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Subscription updated successfully",
		"data":    subscription,
	})
}

func DeleteSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscription models.Subscription
	if result := config.DB.First(&subscription, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Subscription not found",
		})
	}

	config.DB.Delete(&subscription)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Subscription deleted successfully",
	})
}

func GetSubscriptionStats(c *fiber.Ctx) error {
	var totalSubscriptions int64
	var totalAmount float64

	config.DB.Model(&models.Subscription{}).Count(&totalSubscriptions)
	config.DB.Model(&models.Subscription{}).Select("sum(amount)").Row().Scan(&totalAmount)

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"total_subscriptions": totalSubscriptions,
			"total_amount":        totalAmount,
			"monthly_spending":    totalAmount,
		},
	})
}
