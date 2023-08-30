package interfaces

import (
	"context"
	"gochiapp/entities"
	"gochiapp/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Route(r chi.Router)
}

type UserService interface {
	Create(model.CreateUserModel) entities.User
	DeleteOne(id string) string
}

type UserRepository interface {
	Store(entities.User, context.Context)
	Delete(id string, ctx context.Context)
	FindOne(data string, ctx context.Context) (entities.User, error)
}
