package utils

import (
	"time"

	"topinambur02.com/m/v2/model"
)

func FilterByDate(lessons []model.Lesson, targetDate time.Time) []model.Lesson {
	var result []model.Lesson

	for _, lesson := range lessons {
		start, end, err := parseDateInterval(lesson.DateInterval)
		if err != nil {
			continue
		}

		if (targetDate.Equal(start) || targetDate.After(start)) && 
		   (targetDate.Equal(end) || targetDate.Before(end)) {
			result = append(result, lesson)
		}
	}

	return result
}
