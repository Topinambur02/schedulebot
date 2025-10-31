package utils

import (
	"strconv"
	"strings"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
    dateStr = strings.TrimSpace(dateStr)
    parts := strings.Split(dateStr, " ")
    if len(parts) != 2 {
        return time.Time{}, nil
    }

    day, err := strconv.Atoi(parts[0])
    if err != nil {
        return time.Time{}, err
    }

    month := parseMonth(parts[1])
    if month == 0 {
        return time.Time{}, nil
    }

    year := time.Now().Year()
    
    return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), nil
}