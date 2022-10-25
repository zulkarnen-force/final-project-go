package dto

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


type TokenResponse struct {
	Token string `json:"token"`
}


type UserRegisterInput struct {
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"` 
	Password  string 	`json:"password,omitempty"`
	Age       int       `json:"age,omitempty"` 
}

type UserLoginInput struct {
	Email     string    `json:"email,omitempty"` 
	Password  string 	`json:"password,omitempty"`
}