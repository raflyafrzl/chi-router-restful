package user

import (
	"context"
	"errors"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Store(data entities.User, ctx context.Context) {

	err := u.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		panic(err)
	}

}

func (u *userRepository) FindOne(data string, ctx context.Context) (entities.User, error) {

	var user entities.User
	var result *gorm.DB = u.DB.WithContext(ctx).Where("id=? OR email=?", data, data).Find(&user)

	if result.RowsAffected == 0 {
		return entities.User{}, errors.New("No data found")
	}

	return user, nil

}

func (u *userRepository) Delete(id string, ctx context.Context) {

	result := u.DB.WithContext(ctx).Where("id=?", id).Delete(&entities.User{})

	if result.RowsAffected == 0 {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      "Data not found",
		})
	}

}
