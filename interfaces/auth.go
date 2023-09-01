package interfaces

import (
	"gochiapp/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthController interface {
	Route(r chi.Router)
	Login(w http.ResponseWriter, r *http.Request)
}

type AuthService interface {
	CompareAndSigned(data model.LoginUserModel) string
	Set(key string) string
	Get(key string) string
}
