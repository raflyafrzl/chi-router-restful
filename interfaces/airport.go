package interfaces

import (
	"gochiapp/entities"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AirportController interface {
	List(w http.ResponseWriter, r *http.Request)
	Route(r chi.Router)
	Insert(w http.ResponseWriter, r *http.Request)
}
type AirportService interface {
	ListAll() (entities.AirportEntity, error)
}

type AirportRepository interface {
}
