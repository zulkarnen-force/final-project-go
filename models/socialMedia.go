package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)


type SocialMedia struct {
	ID        int `json:"id" gorm:"primarykey"`
	Name  string   `json:"name" valid:"required~Your Name is required"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Your Social Media URL is required"`
	UserID int	 `json:"user_id"`
	User User `json:"user,omitempty" valid:"-"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)

	return
}

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


func (s *SocialMedia) GetResponseCreate() *SocialMediaResponseCreate {
	response := new(SocialMediaResponseCreate)
	copier.Copy(&response, &s)
	return response
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

func (s *SocialMedia) GetResponseUpdate() *SocialMediaResponseUpdate {
	response := new(SocialMediaResponseUpdate)
	copier.Copy(&response, &s)
	return response
}



func (s SocialMedia) ToResponseMedias(socialMediaDatas *[]SocialMedia) *[]SocialMediasResponse {
	var socialMedias []SocialMediasResponse
	
	for _, data := range(*socialMediaDatas) {
		socialMediaResponse := new(SocialMediasResponse)
		copier.Copy(&socialMediaResponse, &data)
		socialMedias = append(socialMedias, *socialMediaResponse)
	}

	return &socialMedias
}