package dto

type ErrorResponse struct {
	Message    string `json:"message"`
	MessageDev string `json:"message_dev,omitempty"`
	Type       string `json:"type,omitempty"`
	Code       int    `json:"code,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}