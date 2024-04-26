package datetime

import "time"

// AddDaysToDate adds the specified number of days to the given date.
func AddDaysToDate(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

// SubtractDaysFromDate subtracts the specified number of days from the given date.
func SubtractDaysFromDate(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

// AddToTime adds the specified years, months, and days to the given date.
func AddToTime(date time.Time, years, months, days int) time.Time {
	return date.AddDate(years, months, days)
}

// SumDates calculates the sum of two dates and returns the result.
func SumDates(date1, date2 time.Time) time.Time {
	return date1.Add(date2.Sub(date1))
}

// GetAge calculates the age based on the given birthdate.
func GetAge(birthdate time.Time) int {
	current := time.Now()
	age := current.Year() - birthdate.Year()

	// Check if the birthday for this year has already occurred.
	if current.YearDay() < birthdate.YearDay() {
		age--
	}
	return age
}
