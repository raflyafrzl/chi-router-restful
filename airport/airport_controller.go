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
}

func (a *airportcontroller) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := a.service.FindAll()
	utils.ErrorResponseWeb(err, 404)
	response, _ := json.Marshal(data)

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

	data, err := a.service.FindAll()
	utils.ErrorResponseWeb(err, 404)

	response, _ := json.Marshal(data)

	w.WriteHeader(201)
	w.Write(response)

}
