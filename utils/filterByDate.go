package utils

import (
	"time"

	"topinambur02.com/m/v2/model"
)

func FilterByDate(lessons []model.Lesson) []model.Lesson {
	var result []model.Lesson
	targetDate := time.Now()

	for _, lesson := range lessons {
		start, end, err := parseDateInterval(lesson.DateInterval)
		if err != nil {
			continue
		}

		if !targetDate.Before(start) && !targetDate.After(end) {
			result = append(result, lesson)
		}
	}

	return result
}
