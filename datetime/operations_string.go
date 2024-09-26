package datetime

// AddDaysStr adds the specified number of days to the given date string and returns the resulting date as a formatted string.
func AddDaysStr(dateStr string, days int) string {
	date := StringToDate(dateStr)
	return DateToString(AddDays(date, days))
}

// SubtractDaysStr subtracts the specified number of days from the given date string and returns the resulting date as a formatted string.
func SubtractDaysStr(dateStr string, days int) string {
	date := StringToDate(dateStr)
	return DateToString(SubtractDays(date, days))
}

// AddWeeksStr adds the specified number of weeks to the given date string and returns the resulting date as a formatted string.
func AddWeeksStr(dateStr string, weeks int) string {
	date := StringToDate(dateStr) 
	return DateToString(AddWeeks(date, weeks))
}

// SubtractWeeksStr subtracts the specified number of weeks from the given date string and returns the resulting date as a formatted string.
func SubtractWeeksStr(dateStr string, weeks int) string {
	date := StringToDate(dateStr)
	return DateToString(SubtractWeeks(date, weeks))
}

// AddMonthsStr adds the specified number of months to the given date string and returns the resulting date as a formatted string.
func AddMonthsStr(dateStr string, months int) string {
	date := StringToDate(dateStr)
	return DateToString(AddMonths(date, months))
}

// SubtractMonthsStr subtracts the specified number of months from the given date string and returns the resulting date as a formatted string.
func SubtractMonthsStr(dateStr string, months int) string {
	date := StringToDate(dateStr)
	return DateToString(SubtractMonths(date, months))
}

// AddYearsStr adds the specified number of years to the given date string and returns the resulting date as a formatted string.
func AddYearsStr(dateStr string, years int) string {
	date := StringToDate(dateStr)
	return DateToString(AddYears(date, years))
}

// SubtractYearsStr subtracts the specified number of years from the given date string and returns the resulting date as a formatted string.
func SubtractYearsStr(dateStr string, years int) string {
	date := StringToDate(dateStr)
	return DateToString(SubtractYears(date, years))
}

// SumDatesStr calculates the sum of two dates in string format and returns the result as a formatted string.
func SumDatesStr(date1Str string, date2Str string) string {
	sum := SumDates(StringToDate(date1Str), StringToDate(date2Str))
	return DateToString(sum)
}

// GetAgeStr calculates the age based on the given birthdate in string format.
func GetAgeStr(birthDateStr string) int {
	birthDate := StringToDate(birthDateStr)
	return GetAge(birthDate)
}
