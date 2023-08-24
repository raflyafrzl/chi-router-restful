package user

import (
	"context"
	"gochiapp/entities"
	"gochiapp/interfaces"

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
