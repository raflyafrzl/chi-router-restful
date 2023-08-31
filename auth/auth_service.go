package auth

import (
	"context"
	"errors"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/redis"
	"gochiapp/utils"
	"math/rand"
	"strconv"
	"time"
)

type authService struct {
	user interfaces.UserRepository
	r    *redis.RedisClient
}

func NewAuthService(user *interfaces.UserRepository, redis *redis.RedisClient) interfaces.AuthService {
	var service authService = authService{
		user: *user,
		r:    redis,
	}

	return &service

}

func (a *authService) Compare(data model.LoginUserModel) {

	utils.Validate[model.LoginUserModel](data)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	var user entities.User
	var err error
	defer cancel()
	user, err = a.user.FindOne(data.Email, ctx)

	utils.ErrorResponseWeb(err, 400)

	if user.Email != data.Email {
		utils.ErrorResponseWeb(errors.New("Invalid email or password"), 400)
	}

	if !utils.VerifyPassword(data.Password, user.Password) {
		utils.ErrorResponseWeb(errors.New("Invalid email or password"), 400)
	}

}

func (a *authService) Set(data string) {

	var randInt *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMicro()))

	var timeLeave = time.Second * 10

	var value string = strconv.Itoa(randInt.Intn(1001) + 9999)

	a.r.SetValue(data, value, timeLeave)

}
