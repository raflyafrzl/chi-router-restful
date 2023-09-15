package interfaces

import (
	"context"
	"gochiapp/entities"
	"gochiapp/model"
)

type UserService interface {
	Create(model.CreateUserModel) entities.User
	DeleteOne(id string) string
	Update(data model.UpdateUserModel, id string) string
}

type UserRepository interface {
	Store(entities.User, context.Context)
	Delete(id string, ctx context.Context)
	FindOne(data string, ctx context.Context) (entities.User, error)
	Update(entities.User, context.Context) error
}
