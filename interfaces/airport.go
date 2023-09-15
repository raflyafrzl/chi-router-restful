package interfaces

import (
	"context"
	"gochiapp/entities"
	"gochiapp/model"
)

type AirportService interface {
	FindAll() ([]entities.Airport, error)
	Create(data model.CreateAirportModel) (model.CreateAirportModel, error)
	FindById(id string) model.CreateAirportModel
	Delete(id string)
	Update(id string, data model.UpdateAirportModel)
}

type AirportRepository interface {
	List(ctx context.Context) ([]entities.Airport, error)
	Store(ctx context.Context, data entities.Airport) (entities.Airport, error)
	Update(ctx context.Context, data entities.Airport) entities.Airport
	First(ctx context.Context, id string) (entities.Airport, error)
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context)
}
