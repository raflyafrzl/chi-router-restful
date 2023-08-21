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

	return s
}

func (s *Server) AllHandlers() {
	s.Router.Use(middlewares.RecoveryMiddleware)
	s.Router.Route("/api/v1/airport", controller.Route)

}

func createTestContext() (context.Context, context.CancelFunc) {

	return context.WithTimeout(context.Background(), time.Second*2)
}

// var app *chi.Mux = StarterServer()
var repository = airport.NewAirportRepository(dbconfig)
var service = airport.NewAirportService(&repository)
var controller = airport.NewAirport(&service)

var configuration = config.New("../.env.test")
var dbconfig = config.InitDB(configuration)
