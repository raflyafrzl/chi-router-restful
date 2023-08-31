package main

import (
	"encoding/json"
	"gochiapp/airport"
	"gochiapp/auth"
	"gochiapp/config"
	"gochiapp/interfaces"
	"gochiapp/middlewares"
	"gochiapp/redis"
	"gochiapp/user"
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

	var userRepository interfaces.UserRepository = user.NewUserRepository(db)
	var userService interfaces.UserService = user.NewUserService(&userRepository)

	var userController interfaces.UserController = user.NewUserController(&userService)

	var redisClient *redis.RedisClient = redis.NewRedisClient(configuration)
	var authService interfaces.AuthService = auth.NewAuthService(&userRepository, redisClient)
	var authController interfaces.AuthController = auth.NewAuthController(&authService)

	//sub-router for airport
	r.Route("/api/v1/airport", airport.Route)
	r.Route("/api/v1/user", userController.Route)
	r.Route("/api/v1/auth", authController.Route)
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
