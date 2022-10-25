package mappers

import (
	"final-project-go/dto"
	"final-project-go/entity"
)

type UserResponseRegister = dto.UserResponseRegister
type UserResponseUpdate = dto.UserResponseUpdate
type UserLoginResponse = dto.TokenResponse
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