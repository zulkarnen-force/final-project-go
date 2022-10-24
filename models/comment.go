package models

import (
	"time"
)

type CommentResponseCreate struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponseUpdate struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	UpdatedAt time.Time `json:"updated_at"`
}


type userResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type photoResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type CommentResponse struct {
	ID        int           `json:"id" gorm:"primarykey"`
	Message   string        `json:"message" valid:"required~Comment is required"`
	PhotoID   int           `json:"photo_id"`
	UserID    int           `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      userResponse  `json:"user"`
	Photo     photoResponse `json:"photo"`
}
