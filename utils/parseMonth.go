package utils

import (
	"strings"
	"time"
)

func parseMonth(monthStr string) time.Month {
	monthMap := map[string]time.Month{
		"Янв": time.January,
		"Фев": time.February,
		"Мар": time.March,
		"Апр": time.April,
		"Май": time.May,
		"Июн": time.June,
		"Июл": time.July,
		"Авг": time.August,
		"Сен": time.September,
		"Окт": time.October,
		"Ноя": time.November,
		"Дек": time.December,
	}

	monthStr = strings.TrimSuffix(monthStr, ".")
	monthStr = strings.TrimSpace(monthStr)

	if month, exists := monthMap[monthStr]; exists {
		return month
	}

	return 0
}
