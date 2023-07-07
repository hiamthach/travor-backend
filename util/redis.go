package util

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/travor-backend/db"
)

type RedisUtil struct {
	redisClient *redis.Client
}

func NewRedisUtil() (*RedisUtil, error) {
	redisClient := db.GetRedisClient()

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
