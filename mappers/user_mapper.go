package mappers

import (
	"final-project-go/entity"
	"final-project-go/models"
)

type UserResponseRegister = models.UserResponseRegister
type UserResponseUpdate = models.UserResponseUpdate
type UserLoginResponse = models.UserLoginResponse
type User = entity.User

func GetResponseRegister(user User) UserResponseRegister {
	return UserResponseRegister{Age:user.Age, ID:user.ID, Username:user.Username, Email:user.Email}
}


func GetResponseUpdate(user User) UserResponseUpdate {
	return UserResponseUpdate{ 
		ID: user.ID,
		Age: user.Age,
		Username: user.Username,
		Email: user.Email,
		UpdatedAt: user.UpdatedAt,
	}
}


func GetUserLoginResponse(tkn string) UserLoginResponse {
	return UserLoginResponse{Token: tkn}
}