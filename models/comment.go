package models

import (
	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int `json:"id" gorm:"primarykey"`
	Message      string    `json:"message" valid:"required~Comment is required"`
	UserID  int   `json:"user_id"`
	PhotoID int `json:"photo_id"`
	User User `json:"user"`
	Photo Photo `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(d *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	return
}