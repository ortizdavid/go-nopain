package datetime

import (
	"fmt"
	"time"
)
// ExtractTimeZone returns the time zone name and offset of the given time
func ExtractTimeZone(date time.Time) (string, int) {
	_, offset := date.Zone()
	return date.Location().String(), offset / 3600
}

// ExtractDateTime returns the date and time in the format YYYY-MM-DD HH:MM:SS
func ExtractDateTime(date time.Time) string {
	return date.Format(dateTimeLayout)
}

// ExtractDate returns the date in the format YYYY-MM-DD
func ExtractDate(date time.Time) string {
	return date.Format(dateLayout)
}

// ExtractYear returns the year as a string in the format YYYY
func ExtractYear(date time.Time) string {
	return date.Format(yearLayout)
}

// ExtractMonth returns the month as a string in the format MM
func ExtractMonth(date time.Time) string {
	return date.Format("01") // Keeping this as is for direct month extraction.
}

// ExtractWeek returns the ISO week number as a string (e.g., "39" for week 39)
func ExtractWeek(date time.Time) string {
	_, week := date.ISOWeek()
	return fmt.Sprintf("%02d", week)
}

// ExtractDay returns the day of the month as a string in the format DD
func ExtractDay(date time.Time) string {
	return date.Format(dayLayout)
}

// ExtractHour returns the time in the format HH:MM:SS
func ExtractHour(date time.Time) string {
	return date.Format(hourLayout)
}

// ExtractSeconds returns only the seconds part (SS) of the time
func ExtractSeconds(date time.Time) string {
	return date.Format(secondsLayout)
}

// ExtractMilliseconds returns the time with milliseconds in the format HH:MM:SS.SSS
func ExtractMilliseconds(date time.Time) string {
	return date.Format(millisecondsLayout)
}