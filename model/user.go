package model

type CreateUserModel struct {
	Name        string `json:"name" validate:"required,min=4"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required"`
}
