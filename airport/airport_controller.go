package airport

import (
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/model"
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

	var data model.ResponseWeb = model.ResponseWeb{
		Message:    "Data Retrieved Successfully",
		StatusCode: 200,
		Data: map[string]interface{}{
			"airport_name": "Icikiwir",
		},
	}

	response, _ := json.Marshal(data)

	w.Write(response)

}

func (a *airportcontroller) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//close body
	defer r.Body.Close()

	var request model.CreateAirportModel

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			Error:      err.Error(),
			StatusCode: 400,
		})
	}

	if err := json.Unmarshal(body, &request); err != nil {
		panic((model.ResponseFailWeb{
			Status:     "Failed",
			Error:      err.Error(),
			StatusCode: 400,
		}))
	}
	data, err := a.service.ListAll()
	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			Error:      err.Error(),
			StatusCode: 400,
		})
	}

	response, _ := json.Marshal(data)

	w.WriteHeader(201)
	w.Write(response)

}
