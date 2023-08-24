package interfaces

import (
	"context"
	"gochiapp/entities"
	"gochiapp/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController interface {
	ListUser(w http.ResponseWriter, r *http.Request)
	Route(r chi.Router)
}

type UserService interface {
	Create(model.CreateUserModel) entities.User
}

type UserRepository interface {
	Store(entities.User, context.Context)
}
