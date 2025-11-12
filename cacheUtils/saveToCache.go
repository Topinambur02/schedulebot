package cacheUtils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"topinambur02.com/m/v2/model"
)

func SaveToCache(lessons []model.Lesson, redisDB *redis.Client, ctx context.Context, key string) error {
	if redisDB == nil {
		return fmt.Errorf("redis client not initialized")
	}

	data, err := json.Marshal(lessons)

	if err != nil {
		return err
	}

	expiration := time.Until(time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour).Add(24 * time.Hour))
	err = redisDB.Set(ctx, key, data, expiration).Err()

	if err != nil {
		return err
	}

	return nil
}
