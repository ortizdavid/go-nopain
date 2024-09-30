package datetime

import (
	"log"
)

// AddDaysStr adds the specified number of days to the given date string and returns the resulting date as a formatted string.
func AddDaysStr(dateStr string, days int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in AddDaysStr: %v", err)
		return "Invalid date"
	}
	return DateToString(AddDays(date, days))
}

// SubtractDaysStr subtracts the specified number of days from the given date string and returns the resulting date as a formatted string.
func SubtractDaysStr(dateStr string, days int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in SubtractDaysStr: %v", err)
		return "Invalid date"
	}
	return DateToString(SubtractDays(date, days))
}

// AddWeeksStr adds the specified number of weeks to the given date string and returns the resulting date as a formatted string.
func AddWeeksStr(dateStr string, weeks int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in AddWeeksStr: %v", err)
		return "Invalid date"
	}
	return DateToString(AddWeeks(date, weeks))
}

// SubtractWeeksStr subtracts the specified number of weeks from the given date string and returns the resulting date as a formatted string.
func SubtractWeeksStr(dateStr string, weeks int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in SubtractWeeksStr: %v", err)
		return "Invalid date"
	}
	return DateToString(SubtractWeeks(date, weeks))
}

// AddMonthsStr adds the specified number of months to the given date string and returns the resulting date as a formatted string.
func AddMonthsStr(dateStr string, months int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in AddMonthsStr: %v", err)
		return "Invalid date"
	}
	return DateToString(AddMonths(date, months))
}

// SubtractMonthsStr subtracts the specified number of months from the given date string and returns the resulting date as a formatted string.
func SubtractMonthsStr(dateStr string, months int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in SubtractMonthsStr: %v", err)
		return "Invalid date"
	}
	return DateToString(SubtractMonths(date, months))
}

// AddYearsStr adds the specified number of years to the given date string and returns the resulting date as a formatted string.
func AddYearsStr(dateStr string, years int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in AddYearsStr: %v", err)
		return "Invalid date"
	}
	return DateToString(AddYears(date, years))
}

// SubtractYearsStr subtracts the specified number of years from the given date string and returns the resulting date as a formatted string.
func SubtractYearsStr(dateStr string, years int) string {
	date, err := StringToDate(dateStr)
	if err != nil {
		log.Printf("Error parsing dateStr in SubtractYearsStr: %v", err)
		return "Invalid date"
	}
	return DateToString(SubtractYears(date, years))
}

// SumDatesStr calculates the sum of two dates in string format and returns the result as a formatted string.
func SumDatesStr(date1Str string, date2Str string) string {
	date1, err1 := StringToDate(date1Str)
	if err1 != nil {
		log.Printf("Error parsing date1Str in SumDatesStr: %v", err1)
		return "Invalid date"
	}
	date2, err2 := StringToDate(date2Str)
	if err2 != nil {
		log.Printf("Error parsing date2Str in SumDatesStr: %v", err2)
		return "Invalid date"
	}
	return DateToString(SumDates(date1, date2))
}

// GetAgeStr calculates the age based on the given birthdate in string format.
func GetAgeStr(birthDateStr string) int {
	birthDate, err := StringToDate(birthDateStr)
	if err != nil {
		log.Printf("Error parsing birthDateStr in GetAgeStr: %v", err)
		return -1
	}
	return GetAge(birthDate)
}
