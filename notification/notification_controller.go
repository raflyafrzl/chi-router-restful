package notification

import (
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type NotificationController struct {
	service interfaces.NotificationService
}

func NewNotificationController(s *interfaces.NotificationService) *NotificationController {

	return &NotificationController{
		service: *s,
	}
}

func (a *NotificationController) Route(r chi.Router) {

	r.Get("/", a.List)

}

func (a *NotificationController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		StatusCode: 201,
		Status:     "Success",
		Message:    "Success created new notification",
		Data:       "",
	}

	var response []byte
	response, _ = json.Marshal(rawResponse)

	w.Write(response)
}

func (a *NotificationController) Delete(w http.ResponseWriter, r *http.Request) {

}

func (a *NotificationController) Update(w http.ResponseWriter, r *http.Request) {

}
func (a *NotificationController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rawResponse model.ResponseWeb = model.ResponseWeb{
		StatusCode: 200,
		Status:     "Success",
		Message:    "Success retrivied notifications",
		Data:       "",
	}

	var response []byte
	response, _ = json.Marshal(rawResponse)

	w.Write(response)
}
