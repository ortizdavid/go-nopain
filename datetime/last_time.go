package datetime

import "time"

// LastDateOfYear returns the last date of the current year.
func LastDateOfYear() time.Time {
	currentYear := time.Now().Year()
	lastDate := time.Date(currentYear, time.December, 31, 0, 0, 0, 0, time.Local)
	return lastDate
}

// LastDayOfCurrentWeek returns the last day (Sunday) of the current week.
func LastDayOfCurrentWeek() time.Time {
	now := time.Now()
	daysUntilSunday := 7 - int(now.Weekday())
	return now.AddDate(0, 0, daysUntilSunday).Truncate(24 * time.Hour).Add(23 * time.Hour + 59 * time.Minute + 59 * time.Second) // Setting time to 23:59:59
}

// LastDayOfCurrentMonth returns the last day of the current month.
func LastDayOfCurrentMonth() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local).Add(-time.Nanosecond) 
}

// LastDayOfMonth returns the last day of the specified month of a given year.
func LastDayOfMonth(year int, month time.Month) time.Time {
	firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
	return firstDayOfNextMonth.Add(-time.Nanosecond) 
}

// LastDayOfWeek returns the last day (Sunday) of the specified week of a given year.
func LastDayOfWeek(year int, week int) time.Time {
	firstDayOfWeek := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, (week-1)*7)
	lastDayOfWeek := firstDayOfWeek.AddDate(0, 0, 6)
	return lastDayOfWeek
}
