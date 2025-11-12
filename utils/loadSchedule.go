package utils

import (
	"context"
	"time"

	"topinambur02.com/m/v2/model"
)

var ctx = context.Background()

func LoadSchedule() ([]model.Lesson, error) {
    today := time.Now()
    return LoadScheduleInternal("schedule:" + today.Format("2006-01-02"), today, today.Weekday())
}
