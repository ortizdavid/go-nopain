package datetime

import "time"

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

func StringToDate(strDate string) time.Time {
	date, _ := time.Parse("2006-01-02", strDate)
	return date
}

func StringToDateTime(strDateTime string) time.Time {
	dateTime, _ := time.Parse("2006-01-02 15:04:05", strDateTime)
	return dateTime
}

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

func DateTimeToString(date time.Time) string {
	return date.Format("2006-01-02")
}

