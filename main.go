package main

import (
	"encoding/json"
	"gochiapp/airport"
	"gochiapp/config"
	"gochiapp/interfaces"
	"gochiapp/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func main() {
	r := chi.NewRouter()

	//setup configuration
	var configuration config.ConfigInf = config.New()
	//settings middleware
	r.Use(middlewares.RecoveryMiddleware)

	//setup repo->service->controller
	var db *gorm.DB = config.InitDB(configuration)
	var airportrepository interfaces.AirportRepository = airport.NewAirportRepository(db)
	var airportService interfaces.AirportService = airport.NewAirportService(&airportrepository)
	var airport interfaces.AirportController = airport.NewAirport(&airportService)
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
	var port string = ":" + configuration.Get("PORT")
	http.ListenAndServe(port, r)

}
