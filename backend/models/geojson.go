package models

import "gorm.io/gorm"

// GeoJSON model to store GeoJSON data in the database
type GeoJSON struct {
	gorm.Model
	Title string `json:"title"`
	Data  string `json:"data"` // GeoJSON data stored as a JSON string
}
