package utils

import (
	"strings"
	"time"
)

func parseDateInterval(interval string) (time.Time, time.Time, error) {
    parts := strings.Split(interval, " - ")
    if len(parts) != 2 {
        return time.Time{}, time.Time{}, nil
    }

    start, err := parseDate(parts[0])
    if err != nil {
        return time.Time{}, time.Time{}, err
    }

    end, err := parseDate(parts[1])
    if err != nil {
        return time.Time{}, time.Time{}, err
    }

    return start, end, nil
}