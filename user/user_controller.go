package user

import (
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type userController struct {
	service interfaces.UserService
}

func NewUserController(service *interfaces.UserService) interfaces.UserController {
	return &userController{
		service: *service,
	}
}

func (u *userController) Route(r chi.Router) {

	r.Post("/", u.ListUser)

}

func (u *userController) ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	utils.ErrorResponseWeb(err, 500)
	var request model.CreateUserModel

	var errorWeb error = json.Unmarshal(body, &request)
	utils.ErrorResponseWeb(errorWeb, 500)

	bodyResponse, err := json.Marshal(request)
	w.WriteHeader(201)
	w.Write(bodyResponse)

}
