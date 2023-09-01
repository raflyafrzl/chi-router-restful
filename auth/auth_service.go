package auth

import (
	"context"
	"errors"
	"gochiapp/config"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/redis"
	"gochiapp/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	user interfaces.UserRepository
	r    *redis.RedisClient
	c    config.ConfigInf
}

func NewAuthService(user *interfaces.UserRepository, redis *redis.RedisClient, c config.ConfigInf) interfaces.AuthService {
	var service authService = authService{
		user: *user,
		r:    redis,
		c:    c,
	}

	return &service

}

func (a *authService) CompareAndSigned(data model.LoginUserModel) string {

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

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":    user.Id,
		"Email": user.Email,
		"Exp":   time.Hour * 1,
	})

	tokenString, err := token.SignedString([]byte(a.c.Get("JWT_KEY")))

	utils.ErrorResponseWeb(err, 401)

	return tokenString

}

func (a *authService) Set(key string) string {

	var randInt *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMicro()))

	var timeLeave = time.Minute * 1

	var value string = strconv.Itoa(randInt.Intn(9999) + 1000)

	data, _ := a.r.GetValue(key)

	if len(data) <= 0 {
		a.r.SetValue(data, value, timeLeave)
		return value
	}
	return data

}
func (a *authService) Get(key string) string {

	data, err := a.r.GetValue(key)

	if err != nil {
		utils.ErrorResponseWeb(errors.New("Invalid OTP"), 404)
	}
	return data

}
