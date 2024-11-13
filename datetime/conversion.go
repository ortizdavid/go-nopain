package datetime

import (
	"fmt"
	"time"
)

// StringToDate converts a string date in the format "2006-01-02" to a time.Time object.
func StringToDate(strDate string) (time.Time, error) {
	date, err := time.Parse(time.DateOnly, strDate)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %v", err)
	}
	return date, nil
}

// StringToDateTime converts a string date and time in the format "2006-01-02 15:04:05" to a time.Time object.
func StringToDateTime(strDateTime string) (time.Time, error) {
	dateTime, err := time.Parse(dateTimeLayout, strDateTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing dateTime: %v", err)
	} 
	return dateTime, nil
}

// DateToString converts a time.Time object to a string date in the format "2006-01-02".
func DateToString(date time.Time) string {
	return date.Format(dateLayout) 
}

// DateTimeToString converts a time.Time object to a string date and time in the format "2006-01-02 15:04:05".
func DateTimeToString(date time.Time) string {
	return date.Format(dateTimeLayout) 
}
