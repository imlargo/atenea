package models

import (
	"time"
)

type Course struct {
	ID        int       `json:"id"`
	Name      string    `binding:"required" json:"name"`
	Email     string    `gorm:"unique" binding:"required" json:"email"`
	Age       uint8     `binding:"required" json:"age"`
	City      *string   `json:"city"` // can be null
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
