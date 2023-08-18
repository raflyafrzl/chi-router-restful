package airport

import (
	"context"
	"errors"
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

func (a *airportrepository) List(ctx context.Context) ([]entities.Airport, error) {
	var result []entities.Airport

	var err error = a.DB.WithContext(ctx).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *airportrepository) Store(ctx context.Context, data entities.Airport) (entities.Airport, error) {
	var err error = a.DB.WithContext(ctx).Create(data).Error

	if err != nil {
		return entities.Airport{}, err
	}

	return data, nil
}

func (a *airportrepository) First(ctx context.Context, id string) (entities.Airport, error) {

	var data entities.Airport

	result := a.DB.WithContext(ctx).Where("id=?", id).First(&data)

	if result.RowsAffected == 0 {
		return entities.Airport{}, errors.New("Airport not found")
	}

	return data, nil

}

func (a *airportrepository) Update(ctx context.Context, data entities.Airport) entities.Airport {

	err := a.DB.WithContext(ctx).Where("id=?", data.Id).Updates(&data).Error
	if err != nil {
		panic(err)
	}

	return data

}

func (a *airportrepository) Delete(ctx context.Context, id string) error {

	err := a.DB.WithContext(ctx).Where("id=?", id).Delete(entities.Airport{}).Error

	if err != nil {
		panic(err)
	}

	return nil

}
