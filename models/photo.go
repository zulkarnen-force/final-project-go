package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Photo struct {
	ID        int `gorm:"primarykey"`
	Title  string   `json:"title" valid="required~Title is required"`
	Caption string `json:"caption"`
	PhotoURL      string    `json:"photo_url" valid="required~Photo URL is required"` 
	UserID int	 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}