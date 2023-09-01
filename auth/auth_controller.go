package auth

import (
	"context"
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/middlewares"
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

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/verify", a.Verify)
	})

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

	var token string = a.service.CompareAndSigned(request)

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Message:    "Login success",
		Data: map[string]string{
			"token": token,
		},
	}

	var response []byte

	response, _ = json.Marshal(rawResponse)

	w.Write(response)

}

func (a *authController) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ctx context.Context = r.Context()
	var rawAuth []byte
	var authData model.UserAuthModel
	rawAuth, _ = json.Marshal(ctx.Value("auth"))

	json.Unmarshal(rawAuth, &authData)

	var otp string = a.service.Set(authData.Id)
	var rawResponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Message:    "OTP Berhasil dikirim",
		Data:       otp,
	}

	var response []byte

	response, _ = json.Marshal(rawResponse)

	w.Write(response)

}
