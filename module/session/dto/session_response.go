package dto

type SessionResponse struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,datavalid"`
	Token    string `json:"token"`
}
