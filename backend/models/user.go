package models

import "gorm.io/gorm"

// GeoJSON model to store GeoJSON data in the database
type User struct {
	gorm.Model
	Email    string `json:"Email"`
	Password string `json:"Password"` // GeoJSON data stored as a JSON string
}
