package controllers

import (
	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/models"
	"github.com/gofiber/fiber/v2"
)

func GetPlans(c *fiber.Ctx) error {
	var plans []models.Plan
	if result := config.DB.Find(&plans); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch plans"})
	}
	return c.JSON(plans)
}

func GetPlan(c *fiber.Ctx) error {
	id := c.Params("id")
	var plan models.Plan
	if result := config.DB.First(&plan, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
	}
	return c.JSON(plan)
}

func CreatePlan(c *fiber.Ctx) error {
	plan := new(models.Plan)
	if err := c.BodyParser(plan); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if result := config.DB.Create(&plan); result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to create plan"})
	}
	return c.Status(201).JSON(plan)
}

func UpdatePlan(c *fiber.Ctx) error {
	id := c.Params("id")
	var plan models.Plan
	if result := config.DB.First(&plan, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
	}

	if err := c.BodyParser(&plan); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	config.DB.Save(&plan)
	return c.JSON(plan)
}

func DeletePlan(c *fiber.Ctx) error {
	id := c.Params("id")
	var plan models.Plan
	if result := config.DB.First(&plan, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Plan not found"})
	}

	config.DB.Delete(&plan)
	return c.JSON(fiber.Map{"message": "Plan deleted successfully"})
}
