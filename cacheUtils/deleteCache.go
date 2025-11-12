package cacheUtils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func DeleteCache(redisDB *redis.Client, ctx context.Context) error {
	_, err := redisDB.FlushDB(ctx).Result()

	if err != nil {
		panic(err)
	}
	
	return nil
}
