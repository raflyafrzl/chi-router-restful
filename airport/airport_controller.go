package airport

import (
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type airportcontroller struct {
	service interfaces.AirportService
}

func NewAirport(s *interfaces.AirportService) interfaces.AirportController {
	return &airportcontroller{service: *s}
}

func (a *airportcontroller) Route(r chi.Router) {
	//TODO: Create All Route for Airport

	r.Get("/", a.List)
	r.Post("/", a.Insert)
	r.Delete("/{id}", a.Delete)
	r.Get("/{id}", a.FindById)
	r.Patch("/{id}", a.Update)
}

func (a *airportcontroller) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := a.service.FindAll()
	utils.ErrorResponseWeb(err, 404)

	var rewresponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		Data:       data,
		StatusCode: 200,
		Message:    "Success retrieved data",
	}

	response, _ := json.Marshal(rewresponse)

	w.Write(response)

}

func (a *airportcontroller) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//close body
	defer r.Body.Close()

	var request model.CreateAirportModel

	body, err := io.ReadAll(r.Body)

	utils.ErrorResponseWeb(err, 400)

	err = json.Unmarshal(body, &request)
	utils.ErrorResponseWeb(err, 400)

	data, err := a.service.Create(request)
	utils.ErrorResponseWeb(err, 404)

	response, _ := json.Marshal(data)

	w.WriteHeader(201)
	w.Write(response)

}

func (a *airportcontroller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var id string = chi.URLParam(r, "id")
	a.service.Delete(id)

	var rawresponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		Message:    "Data has been successfully deleted",
		StatusCode: 200,
		Data:       "id= " + id,
	}

	response, _ := json.Marshal(rawresponse)

	w.Write(response)

}

func (a *airportcontroller) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id string = chi.URLParam(r, "id")
	var data model.CreateAirportModel = a.service.FindById(id)

	var rawresponse model.ResponseWeb = model.ResponseWeb{
		Status:     "Success",
		StatusCode: 200,
		Message:    "Success retrieved a data",
		Data: map[string]interface{}{
			"id":     id,
			"result": data,
		},
	}
	response, _ := json.Marshal(rawresponse)
	w.Write(response)
}

func (a *airportcontroller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	var request model.UpdateAirportModel
	var id string = chi.URLParam(r, "id")
	body, err := io.ReadAll(r.Body)
	utils.ErrorResponseWeb(err, 400)

	////unmarshal body to request
	err = json.Unmarshal(body, &request)
	utils.ErrorResponseWeb(err, 500)

	a.service.Update(id, request)

	w.Write(body)

}
