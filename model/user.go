package model

type CreateUserModel struct {
	Name        string `json:"name" validate:"required,min=4"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=5"`
}

type UpdateUserModel struct {
	Name        string `json:"name" validate:"omitempty,min=4"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,e164"`
	Password    string `json:"password" validate:"omitempty,min=5"`
}

type LoginUserModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
