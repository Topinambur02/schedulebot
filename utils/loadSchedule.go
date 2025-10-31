package utils

import (
	"context"
	"fmt"

	"topinambur02.com/m/v2/cacheUtils"
	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/model"
)

var ctx = context.Background()

func LoadSchedule() ([]model.Lesson, error) {
	redisDB := db.GetClient()
	cachedSchedule, err := cacheUtils.GetFromCache(redisDB, ctx)
	
	if err == nil {
		return cachedSchedule, nil
	}

	schedule, statusCode := GetSchedule()

	if statusCode != 200 {
		return []model.Lesson{}, fmt.Errorf("failed to load schedule, status code: %d", statusCode)
	}

	currentSchedule := GetCurrentSchedule(*schedule)
	filteredLessons := FilterByDate(currentSchedule.Lessons)
	err = cacheUtils.SaveToCache(filteredLessons, redisDB, ctx)

	if err != nil {
		fmt.Printf("Failed to cache schedule: %v\n", err)
	}

	return filteredLessons, nil
}
