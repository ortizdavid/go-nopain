package datetime

import "time"

// LastDateOfYearStr returns the last date of the current year as a formatted string.
func LastDateOfYearStr() string {
	return DateToString(LastDateOfYear())
}

// LastDayOfCurrentWeekStr returns the last day (Sunday) of the current week as a formatted string.
func LastDayOfCurrentWeekStr() string {
	return DateTimeToString(LastDayOfCurrentWeek())
}

// LastDayOfCurrentMonthStr returns the last day of the current month as a formatted string.
func LastDayOfCurrentMonthStr() string {
	return DateToString(LastDayOfCurrentMonth())
}

// LastDayOfMonthStr returns the last day of the specified month and year as a formatted string.
func LastDayOfMonthStr(year int, month int) string {
	return DateToString(LastDayOfMonth(year, time.Month(month)))
}

// LastDayOfWeekStr returns the last day (Sunday) of the specified week and year as a formatted string.
func LastDayOfWeekStr(year int, week int) string {
	return DateToString(LastDayOfWeek(year, week))
}
