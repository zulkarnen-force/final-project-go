package models

import (
	"time"
)

type Comment struct {
	ID        int `gorm:"primarykey"`
	Message      string    `json:"message" valid="required~Comment is required"`
	UserID  int   `json:"user_id"`
	PhotoID string `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}