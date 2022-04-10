package internal

import "time"

func GetStartOfWeek(date time.Time) time.Time {

	weekday := date.Weekday()

	year, month, day := date.Date()
	dateWithDeletedTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return dateWithDeletedTime.Add(-1 * (time.Duration(weekday) - 1) * 24 * time.Hour)
}

func GetEndOfWeek(date time.Time) time.Time {

	weekday := date.Weekday()

	year, month, day := date.Date()
	dateWithDeletedTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	return dateWithDeletedTime.Add((7 - time.Duration(weekday)) * 24 * time.Hour)
}
