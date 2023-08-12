package main

import (
	"encoding/json"
	"gochiapp/airport"
	"gochiapp/config"
	"gochiapp/interfaces"
	"gochiapp/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	//setup configuration
	var config *config.Config = config.New()

	r.Use(middlewares.RecoveryMiddleware)

	var airport interfaces.AirportController = airport.NewAirport()
	//sub-router for airport

	r.Route("/api/v1/airport", airport.Route)
	//if Router not found
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)

		var response map[string]interface{} = map[string]interface{}{
			"error":  "Route not found",
			"status": "failed",
		}

		data, _ := json.Marshal(response)

		w.Write(data)

	})
	var port string = ":" + config.Get("PORT")
	http.ListenAndServe(port, r)

}
