package dto

type RegisterRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=64,containsany=!@#$%^&*(),containsany=0123456789,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type UpdateStaffRequest struct {
	UserID    string `json:"user_id" validate:"required"`
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
