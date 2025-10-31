package utils

import (
	"fmt"

	"topinambur02.com/m/v2/model"
)

func LoadSchedule() ([]model.Lesson, error) {
	schedule, statusCode := GetSchedule()
	
	if statusCode == 200 {
		currentSchedule := GetCurrentSchedule(*schedule)
		return FilterByDate(currentSchedule.Lessons), nil
	}

	return []model.Lesson{}, fmt.Errorf("failed to load schedule, status code: %d", statusCode)
}
