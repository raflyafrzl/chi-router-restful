package test

import (
	"context"
	"gochiapp/airport"
	"gochiapp/config"
	"gochiapp/middlewares"
	"time"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func StarterServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Router.Use(md.RecoveryMiddleware)
	return s
}

func createTestContext() (context.Context, context.CancelFunc) {

	return context.WithTimeout(context.Background(), time.Second*2)
}

// var app *chi.Mux = StarterServer()
var airportRepo = airport.NewAirportRepository(dbconfig)
var airportService = airport.NewAirportService(&airportRepo)
var airportController = airport.NewAirport(&airportService)

var configuration = config.New("../.env.test")
var dbconfig = config.InitDB(configuration)
var md middlewares.Middleware = *middlewares.NewMiddleware(dbconfig)
