package models

import (
	"time"
)


type SocialMedia struct {
	ID        int `gorm:"primarykey"`
	Name  string   `json:"name" valid:"required~Your Name is required"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Your Social Media URL is required"`
	UserID int	 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}