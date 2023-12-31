package auth

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"gochiapp/config"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/redis"
	"gochiapp/utils"
	"html/template"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/gomail.v2"
)

//go:embed mail.htm
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "mail.htm"))

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
		"Name":  user.Name,
		"Exp":   time.Hour * 1,
	})

	tokenString, err := token.SignedString([]byte(a.c.Get("JWT_KEY")))

	utils.ErrorResponseWeb(err, 401)

	return tokenString

}

func (a *authService) Set(authData model.UserAuthModel) {

	var randInt *rand.Rand = rand.New(rand.NewSource(time.Now().UnixMicro()))

	var timeLeave = time.Minute * 5

	var value string = strconv.Itoa(randInt.Intn(9999) + 1000)

	data, _ := a.r.GetValue(authData.Id)

	if len(data) <= 0 {
		a.r.SetValue(authData.Id, value, timeLeave)

		var dataBody map[string]any = map[string]any{
			"Name":  authData.Name,
			"Otp":   value,
			"Brand": "TimeFlies",
		}

		//send to gmail using goroutine
		go a.sendEmail(dataBody, authData.Email)

	}

}
func (a *authService) Get(key string) string {

	data, err := a.r.GetValue(key)

	if err != nil {
		utils.ErrorResponseWeb(errors.New("Invalid OTP"), 404)
	}
	return data
}

func (a *authService) Verified(id string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	var data entities.User = entities.User{
		Id:         id,
		IsVerified: true,
	}

	a.user.Update(data, ctx)

}

func (a *authService) sendEmail(dataBody any, email string) {

	var body bytes.Buffer

	err := myTemplates.Execute(&body, dataBody)

	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", a.c.Get("FROM_EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "TimeFlies Indonesia")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, a.c.Get("FROM_EMAIL"), a.c.Get("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
