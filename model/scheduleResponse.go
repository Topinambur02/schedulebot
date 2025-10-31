package model

type ScheduleResponse struct {
	Friday    DaySchedule `json:"Friday"`
	Monday    DaySchedule `json:"Monday"`
	Saturday  DaySchedule `json:"Saturday"`
	Thursday  DaySchedule `json:"Thursday"`
	Tuesday   DaySchedule `json:"Tuesday"`
	Wednesday DaySchedule `json:"Wednesday"`
}