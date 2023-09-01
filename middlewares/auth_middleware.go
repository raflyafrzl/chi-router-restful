package middlewares

import (
	"context"
	"errors"
	"gochiapp/model"
	"gochiapp/utils"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(han http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string = r.Header.Get("authorization")

		if tokenString == "" {
			utils.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			}

			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			utils.ErrorResponseWeb(errors.New("Invalid token provided"), 401)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			var parentContext context.Context = context.Background()

			var ctx context.Context = context.WithValue(parentContext, "auth", claims)

			requestCtx := r.WithContext(ctx)
			han.ServeHTTP(w, requestCtx)

		} else {

			panic(model.ResponseFailWeb{
				Status:     "Failed",
				StatusCode: 401,
				Error:      err.Error(),
			})
		}

	})

}
