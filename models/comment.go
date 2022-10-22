package models

import (
	"final-project-go/entity"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type User = entity.User

type Comment struct {
	ID        int `json:"id" gorm:"primarykey"`
	Message      string    `json:"message" valid:"required~Comment is required"`
	UserID  int   `json:"user_id"`
	PhotoID int `json:"photo_id"`
	User User `json:"user,omitempty" valid:"-"`
	Photo Photo `json:"photo,omitempty"  valid:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(d *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	return
}


type CommentResponseCreate struct {
	ID        int `json:"id"`
	Message      string    `json:"message"`
	UserID  int   `json:"user_id"`
	PhotoID int `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
}


type CommentResponseUpdate struct {
	ID        int `json:"id"`
	Message      string    `json:"message"`
	UserID  int   `json:"user_id"`
	PhotoID int `json:"photo_id"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (c *Comment) GetResponseCreate() *CommentResponseCreate {
	commentResponse := new(CommentResponseCreate)
	copier.Copy(&commentResponse, &c)
	return commentResponse
}


func (c *Comment) GetResponseUpdate() *CommentResponseUpdate {
	commentResponse := new(CommentResponseUpdate)
	copier.Copy(&commentResponse, &c)
	return commentResponse
}



type userResponse struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type photoResponse struct {
	ID int `json:"id"`
	Title 		string `json:"title"`
	Caption 	string `json:"caption"`
	PhotoURL	string `json:"photo_url"`
	UserID int `json:"user_id"`
}

type CommentResponse struct {
	ID        int `json:"id" gorm:"primarykey"`
	Message      string    `json:"message" valid:"required~Comment is required"`
	PhotoID int `json:"photo_id"`
	UserID  int   `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User userResponse `json:"user"`
	Photo photoResponse `json:"photo"`
}

func (c *Comment) GetResponseComments(commentDatas *[]Comment) *[]CommentResponse {

	var comments []CommentResponse
	for _, data := range(*commentDatas) {
		comment := new(CommentResponse)
		copier.Copy(comment, &data)
		comments = append(comments, *comment)
	}

	return &comments

}