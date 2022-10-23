package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `json:"id,omitempty" gorm:"primarykey"`
	Title     string    `json:"title,omitempty" valid:"required~Title is required"`
	Caption   string    `json:"caption,omitempty"`
	PhotoURL  string    `json:"photo_url,omitempty" valid:"required~Photo URL is required"`
	UserID    int       `json:"user_id,omitempty"`
	User      User      `json:"user,omitempty" valid:"-"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)

	return
}