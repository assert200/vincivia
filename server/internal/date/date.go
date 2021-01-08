package date

import "time"

// CurrDay get current day in unix time
func CurrDay() time.Time {
	return DayOffset(0)
}

// DayOffset returns the day at midnight in unix time
func DayOffset(dayOffset int) time.Time {
	t := time.Now().AddDate(0, 0, dayOffset)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
