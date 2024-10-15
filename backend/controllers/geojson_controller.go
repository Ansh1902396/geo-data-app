package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AnonO6/geo-data-app/services"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type GeoController struct {
	geoService *services.GeoService
}

func NewGeoController(geoService *services.GeoService) *GeoController {
	return &GeoController{geoService}
}

// UploadGeoJSON handles the GeoJSON file upload, validate it, and store it in the database
func (gc *GeoController) UploadGeoJSON(w http.ResponseWriter, r *http.Request) {
	var geoData services.GeoJSONRequest

	// Decode JSON payload
	if err := json.NewDecoder(r.Body).Decode(&geoData); err != nil {
		logrus.Errorf("Failed to decode GeoJSON: %v", err)
		http.Error(w, "Invalid GeoJSON format", http.StatusBadRequest)
		return
	}

	// Pass data to the service for processing
	err := gc.geoService.SaveGeoJSON(&geoData)
	if err != nil {
		logrus.Errorf("Failed to save GeoJSON data: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logrus.Infof("Successfully uploaded GeoJSON with title: %s", geoData.Title)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "GeoJSON uploaded successfully"})
}

// GetGeoJSON retrieves a GeoJSON by ID
func (gc *GeoController) GetGeoJSON(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	geoData, err := gc.geoService.GetGeoJSON(id)
	if err != nil {
		logrus.Errorf("Failed to retrieve GeoJSON with ID %s: %v", id, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	logrus.Infof("Successfully retrieved GeoJSON with ID: %s", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(geoData)
}
