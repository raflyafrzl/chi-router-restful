package middlewares

import (
	"encoding/json"
	"gochiapp/model"
	"net/http"
)

func RecoveryMiddleware(han http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			w.Header().Set("Content-Type", "application/json")
			data := recover()

			if data != nil {
				var err model.ResponseFailWeb = data.(model.ResponseFailWeb)
				w.WriteHeader(err.StatusCode)
				response, _ := json.Marshal(err)
				w.Write(response)
			}

		}()
		han.ServeHTTP(w, r)
	})
}
