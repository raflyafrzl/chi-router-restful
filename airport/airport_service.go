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

func (a *airportservice) FindAll() ([]entities.Airport, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return a.repository.List(ctx)

}

func (a *airportservice) FindById(id string) model.CreateAirportModel {

	var result model.CreateAirportModel
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	data, err := a.repository.First(ctx, id)

	if err != nil {
		utils.ErrorResponseWeb(err, 400)
	}

	result = model.CreateAirportModel{
		Location:        data.Location,
		AirportName:     data.AirportName,
		LocationAcronym: data.LocationAcronym,
	}

	return result
}

func (a *airportservice) Create(data model.CreateAirportModel) (model.CreateAirportModel, error) {
	utils.Validate[model.CreateAirportModel](data)
	var airport entities.Airport

	var airportCode string = string(data.AirportName[0]) + string(data.AirportName[len(data.AirportName)-1]) + "L"
	airport = entities.Airport{
		Id:              utils.GetRandomID(8),
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
func (a airportservice) Delete(id string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	_, err := a.repository.First(ctx, id)

	//not found
	utils.ErrorResponseWeb(err, 404)

	defer cancel()
	a.repository.Delete(ctx, id)

}

func (a *airportservice) Update(id string, data model.UpdateAirportModel) {
	//Validate data
	utils.Validate[model.UpdateAirportModel](data)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	result, err := a.repository.First(ctx, id)
	utils.ErrorResponseWeb(err, 404)

	var airportCode string
	if data.AirportName == "" {
		airportCode = string(result.AirportName[0]) + string(result.AirportName[len(result.AirportName)-1]) + "L"

	} else {
		airportCode = string(data.AirportName[0]) + string(data.AirportName[len(data.AirportName)-1]) + "L"
	}

	var airport entities.Airport = entities.Airport{
		AirportName:     data.AirportName,
		Id:              id,
		Location:        data.Location,
		AirportCode:     airportCode,
		LocationAcronym: data.LocationAcronym,
	}

	a.repository.Update(ctx, airport)

}
