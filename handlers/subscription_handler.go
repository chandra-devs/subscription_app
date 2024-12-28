package handlers

import (
	"fmt"
	"time"

	"github.com/chandra-devs/subscription_app/config"
	"github.com/chandra-devs/subscription_app/models"
	"github.com/gofiber/fiber/v2"
)

// SubscriptionRequest represents the subscription request payload
type SubscriptionRequest struct {
	UserID uint `json:"user_id" validate:"required"`
	PlanID uint `json:"plan_id" validate:"required"`
}

// SubscriptionResponse represents the standardized response
type SubscriptionResponse struct {
	Success bool                 `json:"success"`
	Data    *models.Subscription `json:"data,omitempty"`
	Error   string               `json:"error,omitempty"`
}

// PlanRequest represents the plan request payload
type PlanRequest struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Duration int     `json:"duration" validate:"required"`
}

// PlanResponse represents the standardized response for plans
type PlanResponse struct {
	Success bool         `json:"success"`
	Data    *models.Plan `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
}

func CreatePlan(c *fiber.Ctx) error {
	var req PlanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(PlanResponse{
			Success: false,
			Error:   "Invalid input format",
		})
	}

	plan := models.Plan{
		Name:     req.Name,
		Price:    req.Price,
		Duration: req.Duration,
	}

	if err := config.DB.Create(&plan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(PlanResponse{
			Success: false,
			Error:   "Could not create plan",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(PlanResponse{
		Success: true,
		Data:    &plan,
	})
}

func GetPlans(c *fiber.Ctx) error {
	var plans []models.Plan
	limit := 100 // Or use pagination parameters from request
	if err := config.DB.Limit(limit).Find(&plans).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Could not retrieve plans",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    plans,
	})
}

func GetPlanByID(c *fiber.Ctx) error {
	planID := c.Params("id")
	var plan models.Plan

	if result := config.DB.First(&plan, planID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(PlanResponse{
			Success: false,
			Error:   "Plan not found",
		})
	}

	return c.JSON(PlanResponse{
		Success: true,
		Data:    &plan,
	})
}

// SubscribeUser handles user subscription requests
func SubscribeUser(c *fiber.Ctx) error {
	var req SubscriptionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(SubscriptionResponse{
			Success: false,
			Error:   "Invalid input format",
		})
	}

	// Validate request
	if err := validateSubscriptionRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(SubscriptionResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	// Check if user exists
	var user models.User
	if result := config.DB.First(&user, req.UserID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(SubscriptionResponse{
			Success: false,
			Error:   "User not found",
		})
	}

	// Check if plan exists
	var plan models.Plan
	if result := config.DB.First(&plan, req.PlanID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(SubscriptionResponse{
			Success: false,
			Error:   "Plan not found",
		})
	}

	// Check if user already has an active subscription
	var existingSubscription models.Subscription
	if result := config.DB.Where("user_id = ? AND active = true", req.UserID).First(&existingSubscription); result.Error == nil {
		return c.Status(fiber.StatusBadRequest).JSON(SubscriptionResponse{
			Success: false,
			Error:   "User already has an active subscription",
		})
	}

	subscription := models.Subscription{
		UserID:    req.UserID,
		PlanID:    req.PlanID,
		Status:    "active",
		StartDate: time.Now(),
		ExpiresAt: time.Now().AddDate(0, 0, plan.Duration),
		Active:    true,
	}

	if err := config.DB.Create(&subscription).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(SubscriptionResponse{
			Success: false,
			Error:   "Could not create subscription",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(SubscriptionResponse{
		Success: true,
		Data:    &subscription,
	})
}

func validateSubscriptionRequest(req SubscriptionRequest) error {
	if req.UserID == 0 {
		return fmt.Errorf("user ID is required")
	}
	if req.PlanID == 0 {
		return fmt.Errorf("plan ID is required")
	}
	return nil
}

// GetUserSubscriptions retrieves all subscriptions for a user
func GetUserSubscriptions(c *fiber.Ctx) error {
	userID := c.Params("userId")
	var subscriptions []models.Subscription
	if err := config.DB.Where("user_id = ?", userID).Find(&subscriptions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve subscriptions"})
	}
	return c.JSON(subscriptions)
}

func CreateSubscription(c *fiber.Ctx) error {
	subscription := new(models.Subscription)
	if err := c.BodyParser(subscription); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validate Plan exists
	var plan models.Plan
	if result := config.DB.First(&plan, subscription.PlanID); result.Error != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid plan"})
	}

	// Set subscription details
	subscription.StartDate = time.Now()
	subscription.ExpiresAt = subscription.StartDate.AddDate(0, 0, plan.Duration)
	subscription.Status = "active"
	subscription.Active = true

	if result := config.DB.Create(&subscription); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create subscription"})
	}

	return c.Status(201).JSON(subscription)
}
