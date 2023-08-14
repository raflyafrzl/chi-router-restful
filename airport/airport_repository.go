package airport

import (
	"context"
	"gochiapp/entities"
	"gochiapp/interfaces"

	"gorm.io/gorm"
)

type airportrepository struct {
	DB *gorm.DB
}

func NewAirportRepository(db *gorm.DB) interfaces.AirportRepository {

	return &airportrepository{DB: db}
}

func (a *airportrepository) List(ctx context.Context) (entities.Airport, error) {
	var result entities.Airport

	var err error = a.DB.WithContext(ctx).Find(&result).Error

	if err != nil {
		return entities.Airport{}, err
	}

	return result, nil
}

func (a *airportrepository) Store(ctx context.Context, data entities.Airport) (entities.Airport, error) {
	var err error = a.DB.Create(&data).Error

	if err != nil {
		return entities.Airport{}, err
	}

	return data, nil
}
