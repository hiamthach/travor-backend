package db

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client

	redisHost     = "localhost"
	redisPort     = "6379"
	redisPassword = ""
)

func initRedis(ctx context.Context) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	log.Println("Redis client connected successfully...")
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		initRedis(context.Background())
	}
	return redisClient
}
