package datetime

import (
	"time"
)

// IsValidDateTime checks if a string is a valid date-time in the format "2006-01-02 15:04:05".
func IsValidDateTime(dateTimeStr string) bool {
	_, err := time.Parse(dateTimeLayout, dateTimeStr)
	return err == nil
}

// IsValidDate checks if a string is a valid date in the format "2006-01-02".
func IsValidDate(dateStr string) bool {
	_, err := time.Parse(dateLayout, dateStr)
	return err == nil
}

// IsValidYear checks if a string is a valid year.
func IsValidYear(yearStr string) bool {
	_, err := time.Parse(yearLayout, yearStr)
	return err == nil
}

// IsValidMonth checks if a string is a valid month as part of a full date (YYYY-MM).
func IsValidMonth(monthStr string) bool {
	_, err := time.Parse("2006-01", monthStr)
	return err == nil
}

// IsValidDay checks if a string is a valid day as part of a full date (YYYY-MM-DD).
func IsValidDay(dayStr string) bool {
	_, err := time.Parse(dateLayout, dayStr)
	return err == nil
}

// IsValidHour checks if a string is a valid time in the format "15:04:05".
func IsValidHour(timeStr string) bool {
	_, err := time.Parse("15:04:05", timeStr)
	return err == nil
}

// IsValidSecond checks if a string is a valid time with seconds (SS).
func IsValidSecond(secondStr string) bool {
	_, err := time.Parse("15:04:05", "00:00:"+secondStr)
	return err == nil
}

// IsValidMillisecond checks if a string is a valid time with milliseconds (HH:MM:SS.SSS).
func IsValidMillisecond(timeStr string) bool {
	_, err := time.Parse(millisecondsLayout, "00:00:00."+timeStr)
	return err == nil
}