package datetime

import (
	"time"
)

// StringToDate converts a string date in the format "2006-01-02" to a time.Time object.
func StringToDate(strDate string) time.Time {
	date, _ := time.Parse(dateLayout, strDate)
	return date
}

// StringToDateTime converts a string date and time in the format "2006-01-02 15:04:05" to a time.Time object.
func StringToDateTime(strDateTime string) time.Time {
	dateTime, _ := time.Parse(dateTimeLayout, strDateTime) 
	return dateTime
}

// DateToString converts a time.Time object to a string date in the format "2006-01-02".
func DateToString(date time.Time) string {
	return date.Format(dateLayout) 
}

// DateTimeToString converts a time.Time object to a string date and time in the format "2006-01-02 15:04:05".
func DateTimeToString(date time.Time) string {
	return date.Format(dateTimeLayout) 
}
