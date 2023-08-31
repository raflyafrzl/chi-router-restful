package middlewares

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(han http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string = r.Header.Get("authorization")
		fmt.Println(token)

		han.ServeHTTP(w, r)

	})

}
