package middlewares

import (
	"encoding/json"
	"gochiapp/model"
	"net/http"
)

func (m *Middleware) RecoveryMiddleware(han http.Handler) http.Handler {
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
				w.WriteHeader(500)
				var errResponse map[string]interface{} = map[string]interface{}{
					"Status":     "Error",
					"StatusCode": "500",
					"Error":      data,
				}
				finalresponse, _ := json.Marshal(errResponse)
				w.Write(finalresponse)

			}

		}()
		han.ServeHTTP(w, r)
	})
}
