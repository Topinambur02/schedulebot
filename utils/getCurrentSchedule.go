package utils

import (
	"time"

	"topinambur02.com/m/v2/model"
)

func GetCurrentSchedule(schedule model.ScheduleResponse) model.DaySchedule {
	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		return schedule.Monday
	case time.Tuesday:
		return schedule.Tuesday
	case time.Wednesday:
		return schedule.Wednesday
	case time.Thursday:
		return schedule.Thursday
	case time.Friday:
		return schedule.Friday
	case time.Saturday:
		return schedule.Saturday
	default:
		return model.DaySchedule{}
	}
}
