package datetime

import "time"

// IsValidDate checks if a string is a valid date in the format "2006-01-02".
func IsValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// IsValidDateTime checks if a string is a valid date-time in the format "2006-01-02 15:04:05".
func IsValidDateTime(dateTimeStr string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	return err == nil
}