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

	r.Post("/", u.Create)
	r.Delete("/{id}", u.Delete)

}

func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	utils.ErrorResponseWeb(err, 500)
	var request model.CreateUserModel

	var errorWeb error = json.Unmarshal(body, &request)
	utils.ErrorResponseWeb(errorWeb, 500)

	data := u.service.Create(request)

	bodyResponse, err := json.Marshal(data)
	w.WriteHeader(201)
	w.Write(bodyResponse)

}

func (u *userController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id string = chi.URLParam(r, "id")

	var message string = u.service.DeleteOne(id)

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Message:    message,
	}

	var response []byte
	response, _ = json.Marshal(rawResponse)

	w.WriteHeader(200)
	w.Write(response)

}
