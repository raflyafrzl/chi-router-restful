package main

import (
	"encoding/json"
	"gochiapp/airport"
	"gochiapp/auth"
	"gochiapp/config"
	"gochiapp/interfaces"
	"gochiapp/middlewares"
	"gochiapp/notification"
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
	//setup db connection
	var db *gorm.DB = config.InitDB(configuration)

	//setup Middlewares
	var middlewares middlewares.Middleware = *middlewares.NewMiddleware(db)
	//settings middleware
	r.Use(middlewares.RecoveryMiddleware)

	//Airport
	var airportrepository interfaces.AirportRepository = airport.NewAirportRepository(db)
	var airportService interfaces.AirportService = airport.NewAirportService(&airportrepository)
	var airport *airport.AirportController = airport.NewAirport(&airportService)

	//User
	var userRepository interfaces.UserRepository = user.NewUserRepository(db)
	var userService interfaces.UserService = user.NewUserService(&userRepository)
	var userController *user.UserController = user.NewUserController(&userService, &middlewares)

	var redisClient *redis.RedisClient = redis.NewRedisClient(configuration)

	//Auth
	var authService interfaces.AuthService = auth.NewAuthService(&userRepository, redisClient, configuration)
	var authController *auth.AuthController = auth.NewAuthController(&authService, &middlewares)

	//notification
	var notifRepository interfaces.NotificationRepository = notification.NewNotifRepository(db)
	var notifService interfaces.NotificationService = notification.NewNotifService(&notifRepository)
	var notifController notification.NotificationController = *notification.NewNotificationController(&notifService)

	//sub-router for airport
	r.Route("/api/v1/notification", notifController.Route)
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
