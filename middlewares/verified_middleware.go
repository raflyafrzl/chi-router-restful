package middlewares

import (
	"context"
	"encoding/json"
	"gochiapp/entities"
	"gochiapp/model"
	"net/http"

	"gorm.io/gorm"
)

type Middleware struct {
	db *gorm.DB
}

func NewMiddleware(db *gorm.DB) *Middleware {
	return &Middleware{
		db: db,
	}
}

func (a *Middleware) VerifiedMiddleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var ctx context.Context = r.Context()
		var authData model.UserAuthModel
		var authRaw []byte

		authRaw, _ = json.Marshal(ctx.Value("auth"))
		_ = json.Unmarshal(authRaw, &authData)

		var user entities.User
		a.db.Where("id=?", authData.Id).Find(&user)

		if user.IsVerified == false {
			panic(model.ResponseFailWeb{
				Status:     "Failed",
				StatusCode: 403,
				Error:      "Please verify your account first",
			})
		}

		handler.ServeHTTP(w, r)

	})
}
