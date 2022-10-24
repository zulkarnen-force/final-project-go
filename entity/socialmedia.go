package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int         `json:"id" gorm:"primarykey"`
	Name           string      `json:"name" valid:"required~Your Name is required"`
	SocialMediaURL string      `json:"social_media_url" valid:"required~Your Social Media URL is required"`
	UserID         int         `json:"user_id"`
	User           User `json:"user,omitempty" valid:"-"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)

	return
}