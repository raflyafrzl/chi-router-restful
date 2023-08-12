package airport

import (
	"encoding/json"
	"gochiapp/interfaces"
	"gochiapp/model"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type airport struct{}

func NewAirport() interfaces.AirportController {
	return &airport{}
}

func (a *airport) Route(r chi.Router) {
	//TODO: Create All Route for Airport
	r.Get("/", a.Get)
	r.Post("/", a.Insert)
}

func (a *airport) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := model.ResponseWeb{
		Message:    "Data Retrieved Successfully",
		StatusCode: 200,
		Data: map[string]interface{}{
			"airport_name": "Icikiwir",
		},
	}

	response, _ := json.Marshal(data)

	w.Write(response)

}

func (a *airport) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//close body because r.Body return r.ReadCloser
	defer r.Body.Close()
	var request model.AirportCreateRequest

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

	//FIXME: Change the right request(after get data from DB)

	w.Write([]byte("hello"))

}
