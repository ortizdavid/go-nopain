package datetime

import (
	"log"
	"time"
)

// ExtractDateTimeZoneStr extracts the timezone from a date string.
func ExtractDateTimeZoneStr(dateStr string) string {
	date, err := time.Parse(dateTimeLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractDateTimeZoneStr: %v", err)
		return "Invalid date"
	}
	return date.Location().String()
}

// ExtractDateTimeStr extracts the date and time (YYYY-MM-DD HH:MM:SS).
func ExtractDateTimeStr(dateStr string) string {
	date, err := time.Parse(dateTimeLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractDateTimeStr: %v", err)
		return "Invalid date"
	}
	return date.Format(dateTimeLayout)
}

// ExtractDateStr extracts the date (YYYY-MM-DD).
func ExtractDateStr(dateStr string) string {
	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractDateStr: %v", err)
		return "Invalid date"
	}
	return date.Format(dateLayout)
}

// ExtractYearStr extracts the year.
func ExtractYearStr(dateStr string) string {
	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractYearStr: %v", err)
		return "Invalid date"
	}
	return date.Format("2006")
}

// ExtractMonthStr extracts the month.
func ExtractMonthStr(dateStr string) string {
	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractMonthStr: %v", err)
		return "Invalid date"
	}
	return date.Format("01")
}

// ExtractWeekStr extracts the ISO week.
func ExtractWeekStr(dateStr string) int {
	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractWeekStr: %v", err)
		return -1
	}
	_, week := date.ISOWeek()
	return week
}

// ExtractDayStr extracts the day.
func ExtractDayStr(dateStr string) string {
	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractDayStr: %v", err)
		return "Invalid date"
	}
	return date.Format("02")
}

// ExtractHourStr extracts the hour.
func ExtractHourStr(dateStr string) string {
	date, err := time.Parse(dateTimeLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractHourStr: %v", err)
		return "Invalid date"
	}
	return date.Format("15")
}

// ExtractSecondStr extracts the seconds.
func ExtractSecondStr(dateStr string) string {
	date, err := time.Parse(dateTimeLayout, dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in ExtractSecondStr: %v", err)
		return "Invalid date"
	}
	return date.Format("05")
}
