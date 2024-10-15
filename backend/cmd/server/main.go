package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AnonO6/geo-data-app/config"
	"github.com/AnonO6/geo-data-app/controllers"
	"github.com/AnonO6/geo-data-app/middleware"
	"github.com/AnonO6/geo-data-app/models"
	"github.com/AnonO6/geo-data-app/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	// Initialize config, DB, and Redis connection
	config.LoadEnv()
	db := config.InitDB()
	db.AutoMigrate(&models.User{}, &models.GeoJSON{})
	redisClient := config.InitRedis()

	// Initialize services
	authService := services.NewAuthService(db, redisClient)
	geoService := services.NewGeoService(db, redisClient)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	geoController := controllers.NewGeoController(geoService)

	// Setup Router
	r := mux.NewRouter()

	// Auth Routes
	r.HandleFunc("/api/register", authController.Register).Methods("POST")
	r.HandleFunc("/api/login", authController.Login).Methods("POST")

	// Protected Routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.HandleFunc("/geojson/upload", geoController.UploadGeoJSON).Methods("POST")
	api.HandleFunc("/geojson/{id}", geoController.GetGeoJSON).Methods("GET")
	// api.HandleFunc("/geojson/update/{id}", geoController.UpdateGeoJSON).Methods("PUT")

	// Start the server
	port := os.Getenv("PORT")
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
