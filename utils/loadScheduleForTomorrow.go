package utils

import (
	"time"

	"topinambur02.com/m/v2/model"
)

func LoadScheduleForTomorrow() ([]model.Lesson, error) {
    tomorrow := time.Now().Add(24 * time.Hour)
    return LoadScheduleInternal("tomorrow", tomorrow, tomorrow.Weekday())
}