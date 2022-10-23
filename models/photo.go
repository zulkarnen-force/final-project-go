package models

import (
	"final-project-go/entity"
	"time"
)

type Photo = entity.Photo


type ResponsePhoto struct {
	ID int `json:"id"`
	Title 		string `json:"title"`
	Caption 	string `json:"caption"`
	PhotoURL	string `json:"photo_url"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoResponseUpdate struct {
	ID int `json:"id"`
	Title 		string `json:"title"`
	Caption 	string `json:"caption"`
	PhotoURL	string `json:"photo_url"`
	UserID	int `json:"user_id"`
	UpdatedAt	time.Time `json:"updated_at"`
}


type ResponseGetPhotos struct {
	ID        int `json:"id,omitempty" gorm:"primarykey"`
	Title  string   `json:"title,omitempty" valid:"required~Title is required"`
	Caption string `json:"caption,omitempty"`
	PhotoURL      string    `json:"photo_url,omitempty" valid:"required~Photo URL is required"` 
	UserID int	 `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User struct{
		Username string `json:"username"`
		Email string `json:"email"`
		} `json:"user,omitempty"`
}
