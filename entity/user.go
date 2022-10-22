package entity

import (
	"final-project-go/helpers"
	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primarykey" json:"id,omitempty"`
	Username  string    `json:"username,omitempty" valid:"required~Your Username is required"`
	Email     string    `json:"email,omitempty" valid:"email,required~Your Email is required"`
	Password  string 	`json:"password,omitempty" valid:"required~Your Password is required,minstringlength(6)~Password has to have a minumum length of 6 characters"`
	Age       int       `json:"age,omitempty" valid:"required~Your Age is required,range(9|255)~Minimum age must be more than 8"` 
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil 
	return
}