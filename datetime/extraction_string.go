package datetime

import (
	"time"
)


// ExtractDateTimeZone extracts the timezone from a date string.
func ExtractDateTimeZone(dateStr string) string {
	date, _ := time.Parse(dateTimeLayout, dateStr)
	return date.Location().String()
}

// ExtractDateTimeStr extracts the date and time (YYYY-MM-DD HH:MM:SS).
func ExtractDateTimeStr(dateStr string) string {
	date, _ := time.Parse(dateTimeLayout, dateStr)
	return date.Format(dateTimeLayout)
}

// ExtractDateStr extracts the date (YYYY-MM-DD).
func ExtractDateStr(dateStr string) string {
	date, _ := time.Parse(dateLayout, dateStr)
	return date.Format(dateLayout)
}

// ExtractYearStr extracts the year.
func ExtractYearStr(dateStr string) string {
	date, _ := time.Parse(dateLayout, dateStr)
	return date.Format("2006")
}

// ExtractMonthStr extracts the month.
func ExtractMonthStr(dateStr string) string {
	date, _ := time.Parse(dateLayout, dateStr)
	return date.Format("01")
}

// ExtractWeekStr extracts the ISO week.
func ExtractWeekStr(dateStr string) int {
	date, _ := time.Parse(dateLayout, dateStr)
	_, week := date.ISOWeek()
	return week
}

// ExtractDayStr extracts the day.
func ExtractDayStr(dateStr string) string {
	date, _ := time.Parse(dateLayout, dateStr)
	return date.Format("02")
}

// ExtractHourStr extracts the hour.
func ExtractHourStr(dateStr string) string {
	date, _ := time.Parse(dateTimeLayout, dateStr)
	return date.Format("15")
}

// ExtractSecondStr extracts the seconds.
func ExtractSecondStr(dateStr string) string {
	date, _ := time.Parse(dateTimeLayout, dateStr)
	return date.Format("05")
}