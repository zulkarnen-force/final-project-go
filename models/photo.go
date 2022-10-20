package models

import (
	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type Photo struct {
	ID        int `json:"id,omitempty" gorm:"primarykey"`
	Title  string   `json:"title,omitempty" valid:"required~Title is required"`
	Caption string `json:"caption,omitempty"`
	PhotoURL      string    `json:"photo_url,omitempty" valid:"required~Photo URL is required"` 
	UserID int	 `json:"user_id,omitempty"`
	User User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}


func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)

	return
}

type ResponsePhoto struct {
	ID int `json:"id"`
	Title 		string `json:"title"`
	Caption 	string `json:"caption"`
	PhotoURL	string `json:"photo_url"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}


func (p *Photo) ResponsePhotoCreate() ResponsePhoto {
	return ResponsePhoto{
		ID: p.ID,
		Title: p.Title,
		Caption: p.Caption,
		PhotoURL: p.PhotoURL,
		UserID: p.UserID,
		CreatedAt: p.CreatedAt,
	}
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


func (p *Photo) ResponseGetPhotos(photosData *[]Photo) *[]ResponseGetPhotos {
	
	var photos []ResponseGetPhotos = []ResponseGetPhotos{}

	photo := ResponseGetPhotos{}
	
	for _, data := range(*photosData) {
		photo.ID = data.ID
		photo.Title = data.Title
		photo.Caption = data.Caption
		photo.PhotoURL = data.PhotoURL
		photo.UserID = data.UserID
		photo.CreatedAt = data.CreatedAt
		photo.UpdatedAt = data.UpdatedAt

		photo.User.Email = data.User.Email
		photo.User.Username = data.User.Username

		photos = append(photos, photo)
	}

	return &photos
	
}