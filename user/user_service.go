package user

import (
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/utils"
)

type userService struct {
	repository interfaces.UserRepository
}

func NewUserService(repository *interfaces.UserRepository) interfaces.UserService {
	return &userService{
		repository: *repository,
	}
}

func (u *userService) Create(data model.CreateUserModel) entities.User {

	utils.Validate[model.CreateUserModel](data)

	var result entities.User = entities.User{
		Id:          utils.GetRandomID(12),
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		Role:        "user",
		PhoneNumber: data.PhoneNumber,
	}

	return result

}
