package cacheUtils

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func DeleteCache(redisDB *redis.Client, ctx context.Context) error {
	key := "schedule:" + time.Now().Format("2006-01-02")
	_, err := redisDB.Del(ctx, key).Result()

	if err != nil {
		panic(err)
	}
	
	return nil
}
