package interfaces

import (
	"context"
	"gochiapp/entities"
	"gochiapp/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AirportController interface {
	List(w http.ResponseWriter, r *http.Request)
	Route(r chi.Router)
	Insert(w http.ResponseWriter, r *http.Request)
}
type AirportService interface {
	FindAll() (entities.Airport, error)
	Create(data model.CreateAirportModel) (model.CreateAirportModel, error)
}

type AirportRepository interface {
	List(ctx context.Context) (entities.Airport, error)
	Store(ctx context.Context, data entities.Airport) (entities.Airport, error)
}
