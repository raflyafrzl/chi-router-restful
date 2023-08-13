package airport

import (
	"gochiapp/entities"
	"gochiapp/interfaces"
)

type airportservice struct{}

func NewAirportService() interfaces.AirportService {
	return &airportservice{}
}

func (a *airportservice) ListAll() (entities.AirportEntity, error) {
	//FIXME: Change to actual logic
	var data entities.AirportEntity = entities.AirportEntity{
		Id:          "Rafly",
		AirportName: "TestAirport",
		AirportCode: "LCK",
	}

	return data, nil

}
