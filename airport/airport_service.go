package airport

import (
	"context"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/utils"
	"time"
)

type airportservice struct {
	repository interfaces.AirportRepository
}

func NewAirportService(r *interfaces.AirportRepository) interfaces.AirportService {
	return &airportservice{repository: *r}
}

func (a *airportservice) FindAll() (entities.Airport, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return a.repository.List(ctx)

}

func (a *airportservice) Create(data model.CreateAirportModel) (model.CreateAirportModel, error) {
	utils.Validate[model.CreateAirportModel](data)
	var airport entities.Airport

	var airportCode string = string(data.AirportName[0]) + string(data.AirportName[len(data.AirportName)-1]) + "L"
	airport = entities.Airport{
		Id:              utils.GetRandomID(),
		AirportName:     data.AirportName,
		Location:        data.Location,
		AirportCode:     airportCode,
		LocationAcronym: data.LocationAcronym,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	_, err := a.repository.Store(ctx, airport)
	if err != nil {
		return model.CreateAirportModel{}, err
	}
	return data, nil

}
