package config

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

func LoadEnv() {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "postgresql://geodrawdb_owner:LTv4gunpcx1B@ep-dry-cherry-a128tdxu.ap-southeast-1.aws.neon.tech/geodrawdb?sslmode=require")
	os.Setenv("REDIS_URL", "localhost:6379")
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	return db
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	return client
}
