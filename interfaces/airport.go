package interfaces

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AirportController interface {
	Get(w http.ResponseWriter, r *http.Request)
	Route(r chi.Router)
	Insert(w http.ResponseWriter, r *http.Request)
}
