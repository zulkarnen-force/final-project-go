package dto

import (
	"time"
)


type SocialMediaResponseCreate struct {
	ID        int `json:"id" gorm:"primarykey"`
	Name  string   `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID int	 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}


type SocialMediaResponseUpdate struct {
	ID        int `json:"id" gorm:"primarykey"`
	Name  string   `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID int	 `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}


type SocialMediasResponse struct {
	ID        int `json:"id" gorm:"primarykey"`
	Name  string   `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID int	 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User userResponse `json:"user"`
}
