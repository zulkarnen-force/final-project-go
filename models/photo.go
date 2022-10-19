package models

import (
	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type Photo struct {
	ID        int `json:"id" gorm:"primarykey"`
	Title  string   `json:"title" valid:"required~Title is required"`
	Caption string `json:"caption"`
	PhotoURL      string    `json:"photo_url" valid:"required~Photo URL is required"` 
	UserID int	 `json:"user_id"`
	User User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)

	return
}