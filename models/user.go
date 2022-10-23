package models

import (
	"time"
)

type UserResponseRegister struct {
	Age int `json:"age"`
	ID int`json:"id"`
	Username string `json:"username"`
	Email string `json:"email,omitempty"`
}


type UserResponseUpdate struct {
	Age int `json:"age"`
	ID int`json:"id"`
	Username string `json:"username"`
	Email string `json:"email,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}


type UserLoginResponse struct {
	Token string `json:"token"`
}
