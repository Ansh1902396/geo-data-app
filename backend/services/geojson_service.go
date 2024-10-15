package services

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/AnonO6/geo-data-app/models"
	"github.com/go-redis/redis/v8"
	geojson "github.com/paulmach/go.geojson"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GeoService struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewGeoService(db *gorm.DB, redisClient *redis.Client) *GeoService {
	return &GeoService{db, redisClient}
}

type GeoJSONRequest struct {
	Title string                     `json:"title"`
	Data  *geojson.FeatureCollection `json:"data"`
}

// SaveGeoJSON validates and stores the GeoJSON in the database
func (gs *GeoService) SaveGeoJSON(req *GeoJSONRequest) error {
	// Validate GeoJSON format using the go.geojson library
	if req.Data == nil {
		logrus.Error("Invalid GeoJSON format, data is nil")
		return errors.New("invalid GeoJSON data")
	}

	// Convert the GeoJSON data to JSON string
	geojsonBytes, err := json.Marshal(req.Data)
	if err != nil {
		logrus.Errorf("Error marshaling GeoJSON data: %v", err)
		return err
	}

	// Store the GeoJSON data in the database
	geojsonModel := models.GeoJSON{
		Title: req.Title,
		Data:  string(geojsonBytes),
	}

	if err := gs.db.Create(&geojsonModel).Error; err != nil {
		logrus.Errorf("Error saving GeoJSON to database: %v", err)
		return err
	}

	logrus.Infof("GeoJSON successfully saved with title: %s", req.Title)
	return nil
}

// GetGeoJSON retrieves GeoJSON from the cache or database
func (gs *GeoService) GetGeoJSON(id string) (*models.GeoJSON, error) {
	ctx := context.Background()

	// Try to get GeoJSON from Redis cache
	cachedData, err := gs.redisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		logrus.Infof("Cache miss for GeoJSON with ID: %s", id)
	} else if err != nil {
		logrus.Errorf("Error retrieving GeoJSON from cache: %v", err)
		return nil, err
	} else {
		// Cache hit, unmarshal and return
		var geojsonModel models.GeoJSON
		if err := json.Unmarshal([]byte(cachedData), &geojsonModel); err != nil {
			logrus.Errorf("Error unmarshaling cached GeoJSON: %v", err)
			return nil, err
		}

		logrus.Infof("Cache hit for GeoJSON with ID: %s", id)
		return &geojsonModel, nil
	}

	// Cache miss, fetch from database
	var geojsonModel models.GeoJSON
	if err := gs.db.First(&geojsonModel, "id = ?", id).Error; err != nil {
		logrus.Errorf("Error fetching GeoJSON from database: %v", err)
		return nil, errors.New("GeoJSON not found")
	}

	// Cache the result in Redis with a TTL of 30 minutes
	geojsonBytes, _ := json.Marshal(geojsonModel)
	err = gs.redisClient.Set(ctx, id, geojsonBytes, 30*time.Minute).Err()
	if err != nil {
		logrus.Errorf("Error caching GeoJSON in Redis: %v", err)
	}

	logrus.Infof("Successfully retrieved GeoJSON with ID: %s from database", id)
	return &geojsonModel, nil
}
