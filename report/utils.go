package report

import "time"

func getToday() time.Time {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	return today
}

func getYesterday() time.Time {
	yesterday := getToday().AddDate(0, 0, -1)

	return yesterday
}
