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
				err, ok := data.(model.ResponseFailWeb)

				if ok {
					w.WriteHeader(err.StatusCode)
					response, _ := json.Marshal(err)
					w.Write(response)
					return
				}

			}

		}()
		han.ServeHTTP(w, r)
	})
}
