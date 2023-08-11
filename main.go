package main

import (
	"encoding/json"
	"gochiapp/airport"
	"gochiapp/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Message string `json:"message"`
	Status  int16  `json:"status"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middlewares.RecoveryMiddleware)

	airport := airport.NewAirport()
	//sub-router for airport
	r.Route("/api/v1/airport", airport.Route)
	//if Router not found
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)

		response := map[string]interface{}{
			"error":  "Route not found",
			"status": "failed",
		}

		data, _ := json.Marshal(response)

		w.Write(data)

	})

	http.ListenAndServe(":3000", r)

}
