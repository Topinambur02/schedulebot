package utils

import (
	"fmt"
	"strings"
	"time"
)

func parseDateInterval(interval string) (time.Time, time.Time, error) {
	parts := strings.Split(interval, " - ")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid interval format: %s", interval)
	}

	start, err := parseDate(parts[0])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	end, err := parseDate(parts[1])
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if end.Before(start) {
		end = end.AddDate(1, 0, 0)
	}

	return start, end, nil
}
