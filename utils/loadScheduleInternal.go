package utils

import (
	"fmt"
	"time"

	"topinambur02.com/m/v2/cacheUtils"
	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/model"
)

func LoadScheduleInternal(cacheKey string, targetDate time.Time, weekday time.Weekday) ([]model.Lesson, error) {
    redisDB := db.GetClient()
    cached, err := cacheUtils.GetFromCache(redisDB, ctx, cacheKey)

    if err == nil {
        return cached, nil
    }

    schedule, statusCode := GetSchedule()
	
    if statusCode != 200 {
        return nil, fmt.Errorf("failed to load schedule, status code: %d", statusCode)
    }

    currentSchedule := GetScheduleByWeekDay(*schedule, weekday)
    filteredLessons := FilterByDate(currentSchedule.Lessons, targetDate)

    if err := cacheUtils.SaveToCache(filteredLessons, redisDB, ctx, cacheKey); err != nil {
        fmt.Printf("Failed to cache schedule: %v\n", err)
    }

    return filteredLessons, nil
}