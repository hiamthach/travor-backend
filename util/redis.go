package util

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisUtil struct {
	redisClient *redis.Client
}

func NewRedisUtil(config Config) (*RedisUtil, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: config.RedisPassword,
		Username: config.RedisUsername,
		DB:       0,
	})

	log.Println("Redis client connected successfully...")

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisUtil{redisClient: redisClient}, nil
}

func (util *RedisUtil) Get(ctx context.Context, key string) (string, error) {
	return util.redisClient.Get(ctx, key).Result()
}

func (util *RedisUtil) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return util.redisClient.Set(ctx, key, value, expiration).Err()
}

func (util *RedisUtil) Clear(ctx context.Context, key string) error {
	return util.redisClient.Del(ctx, key).Err()
}

func (util *RedisUtil) ClearWithPrefix(ctx context.Context, prefix string) error {
	keys, err := util.redisClient.Keys(ctx, prefix+"*").Result()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	return util.redisClient.Del(ctx, keys...).Err()
}

func (util *RedisUtil) ClearAll() error {
	ctx := context.Background()
	return util.redisClient.FlushAll(ctx).Err()
}
