package auth

import (
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/model"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type authController struct {
	service interfaces.AuthService
}

func NewAuthController(s *interfaces.AuthService) interfaces.AuthController {

	return &authController{
		service: *s,
	}

}

func (a *authController) Route(r chi.Router) {

	r.Post("/login", a.Login)
	r.Post("/verify/{id}", a.Verify)

}

func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request model.LoginUserModel

	rawBody, error := io.ReadAll(r.Body)

	if error != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			Error:      "Invalid request payload",
			StatusCode: 400,
		})
	}
	//unmarshal data
	json.Unmarshal(rawBody, &request)

	a.service.Compare(request)

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Message:    "Login success",
		Data:       "token",
	}

	var response []byte

	response, _ = json.Marshal(rawResponse)

	w.Write(response)

}

func (a *authController) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Message:    "Login success",
		Data:       "token",
	}

	var id string = chi.URLParam(r, "id")

	a.service.Set(id)

	var response []byte

	response, _ = json.Marshal(rawResponse)

	w.Write(response)

}
