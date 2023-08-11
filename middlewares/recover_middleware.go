package middlewares

import "net/http"

func recoveryMiddleware(han http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				w.Write([]byte("Test"))
			}

		}()
		han.ServeHTTP(w, r)
	})
}
