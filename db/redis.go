package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func InitRedisClient() *redis.Client {
	redisOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file, using system environment variables")
		}

		addr := os.Getenv("REDIS_ADDR")
		options := &redis.Options{Addr: addr}
		redisClient = redis.NewClient(options)
		ctx := context.Background()
		_, err := redisClient.Ping(ctx).Result()

		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}

		log.Println("Connected to Redis!")
	})

	return redisClient
}

func GetClient() *redis.Client {
	if redisClient == nil {
		return InitRedisClient()
	}
	return redisClient
}
