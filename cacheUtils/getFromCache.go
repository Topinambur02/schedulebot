package cacheUtils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"topinambur02.com/m/v2/model"
)

func GetFromCache(redisDB *redis.Client, ctx context.Context) ([]model.Lesson, error) {
	if redisDB == nil {
		return nil, fmt.Errorf("redis client not initialized")
	}

	data, err := redisDB.Get(ctx, "schedule:" + time.Now().Format("2006-01-02")).Result()

	if err != nil {
		return nil, err
	}

	var lessons []model.Lesson
	err = json.Unmarshal([]byte(data), &lessons)

	if err != nil {
		return nil, err
	}

	return lessons, nil
}
