package datetime

import "time"

// Today returns the current date and time.
func Today() time.Time {
	return time.Now()
}

// CurrentDate returns the current date in the format "2006-01-02".
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// CurrentDateTime returns the current date and time in the format "2006-01-02 15:04:05".
func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurrentYear returns the current year.
func CurrentYear() int {
	currentTime := time.Now()
	return currentTime.Year()
}

// LastDateOfYear returns the last date of the current year.
func LastDateOfYear() time.Time {
	currentYear := time.Now().Year()
	lastDateOfYear := time.Date(currentYear, time.December, 31, 0, 0, 0, 0, time.Local)
	return lastDateOfYear
}
