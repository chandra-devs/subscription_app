// models/user.go
package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
// @Description User account information
type User struct {
	// Standard fields from gorm.Model
	ID        uint           `json:"id" gorm:"primarykey" example:"1"`
	CreatedAt time.Time      `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time"`

	// User specific fields
	Name     string `json:"name" gorm:"size:255;not null" example:"John Doe"`
	Email    string `json:"email" gorm:"size:255;not null;unique" example:"john@example.com"`
	Password string `json:"-" gorm:"size:255;not null"` // Password is not exposed in JSON

	// Relationships
	Subscriptions []Subscription `json:"subscriptions,omitempty" gorm:"foreignKey:UserID"`
}
