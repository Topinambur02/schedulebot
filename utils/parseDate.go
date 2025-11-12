package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	parts := strings.Split(dateStr, " ")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid day: %s", parts[0])
	}

	month := parseMonth(parts[1])
	if month == 0 {
		return time.Time{}, fmt.Errorf("invalid month: %s", parts[1])
	}

	year := time.Now().Year()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}
