package user

import (
	"context"
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/middlewares"
	"gochiapp/model"
	"gochiapp/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	service interfaces.UserService
	m       middlewares.Middleware
}

func NewUserController(service *interfaces.UserService, m *middlewares.Middleware) *UserController {
	return &UserController{
		service: *service,
		m:       *m,
	}
}

func (u *UserController) Route(r chi.Router) {

	r.Post("/", u.Create)
	r.Delete("/{id}", u.Delete)

	r.Group(func(r chi.Router) {
		r.Use(u.m.AuthMiddleware)
		r.Use(u.m.VerifiedMiddleware)
		r.Patch("/", u.Update)
	})

}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
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

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
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

func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request model.UpdateUserModel
	var authData model.UserAuthModel
	var rawAuth []byte

	var ctx context.Context = r.Context()

	rawAuth, _ = json.Marshal(ctx.Value("auth"))

	_ = json.Unmarshal(rawAuth, &authData)

	rawBody, err := io.ReadAll(r.Body)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			Error:      "Invalid payload request",
			StatusCode: 400,
		})
	}

	_ = json.Unmarshal(rawBody, &request)

	var message string = u.service.Update(request, authData.Id)

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Data:       authData.Id,
		Message:    message,
	}
	var response []byte
	response, _ = json.Marshal(rawResponse)

	w.WriteHeader(200)
	w.Write(response)

}
