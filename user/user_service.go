package user

import (
	"context"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type userService struct {
	repository interfaces.UserRepository
}

func NewUserService(repository *interfaces.UserRepository) interfaces.UserService {

	validate = validator.New()

	return &userService{
		repository: *repository,
	}
}

func (u *userService) Create(data model.CreateUserModel) entities.User {

	utils.Validate[model.CreateUserModel](data)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	hashedPassword, err := utils.HashPassword(data.Password)

	utils.ErrorResponseWeb(err, 500)

	var result entities.User = entities.User{
		Id:          utils.GetRandomID(12),
		Name:        data.Name,
		Email:       data.Email,
		Password:    string(hashedPassword),
		Role:        "user",
		PhoneNumber: data.PhoneNumber,
	}

	u.repository.Store(result, ctx)

	return result

}

func (u *userService) DeleteOne(id string) string {
	var error error
	error = validate.Var(id, "email")

	if error == nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      "Invalid Id Provided",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	_, error = u.repository.FindOne(id, ctx)

	if error != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 404,
			Error:      error.Error(),
		})
	}

	u.repository.Delete(id, ctx)
	defer cancel()

	return "User has been successfully deleted"

}

func (u *userService) Update(data model.UpdateUserModel, id string) string {

	utils.Validate[model.UpdateUserModel](data)

	var error error

	error = validate.Var(id, "email")

	if error == nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      "Invalid id parameter",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	_, error = u.repository.FindOne(id, ctx)

	if error != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			Error:      error.Error(),
			StatusCode: 404,
		})
	}

	defer cancel()

	var user entities.User = entities.User{
		Id:          id,
		Name:        data.Name,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	}

	error = u.repository.Update(user, ctx)

	if error != nil {
		panic(model.ResponseFailWeb{
			Status:     "Error",
			Error:      error.Error(),
			StatusCode: 500,
		})
	}

	return "Data has been successfully updated"

}
