package datetime

import "time"


func Today() time.Time {
	return time.Now()
}


func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}


func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}


func CurrentYear() int {
	currentTime := time.Now()
	return currentTime.Year()
}


func LastDateOfYear() time.Time {
	currentYear := time.Now().Year()
	lastDateOfYear := time.Date(currentYear, time.December, 31, 0, 0, 0, 0, time.Local)
	return lastDateOfYear
}