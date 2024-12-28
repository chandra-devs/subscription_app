// models/subscription.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// Subscription represents the subscription model
// @Description Subscription information
type Subscription struct {
	// Standard fields from gorm.Model
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time"`

	// Subscription specific fields
	UserID    uint      `json:"user_id" gorm:"not null" example:"1" validate:"required"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	PlanID    uint      `json:"plan_id" gorm:"not null" example:"1" validate:"required"`
	Plan      Plan      `json:"plan,omitempty" gorm:"foreignKey:PlanID"`
	Status    string    `json:"status" gorm:"size:50;not null" example:"active" validate:"required,oneof=active expired cancelled"`
	StartDate time.Time `json:"start_date" example:"2024-01-01T00:00:00Z"`
	ExpiresAt time.Time `json:"expires_at" example:"2024-02-01T00:00:00Z"`
	Active    bool      `json:"active" gorm:"default:true" example:"true"`
}

// Plan represents the subscription plan model
// @Description Plan information
type Plan struct {
	// Standard fields from gorm.Model
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time"`

	// Plan specific fields
	Name        string  `json:"name" gorm:"size:255;not null;unique" example:"Premium Plan"`
	Description string  `json:"description" gorm:"size:1000" example:"Premium features included"`
	Price       float64 `json:"price" gorm:"not null" example:"29.99"`
	Duration    int     `json:"duration" gorm:"not null" example:"30"` // duration in days
}
