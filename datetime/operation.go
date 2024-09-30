package datetime

import "time"

// AddDays adds the specified number of days to the given date.
func AddDays(date time.Time, days int) time.Time {
    return date.AddDate(0, 0, days)
}

// SubtractDays subtracts the specified number of days from the given date.
func SubtractDays(date time.Time, days int) time.Time {
    return date.AddDate(0, 0, -days)
}

// AddWeeks adds the specified number of weeks to the given date.
func AddWeeks(date time.Time, weeks int) time.Time {
    return date.AddDate(0, 0, weeks*7) 
}

// SubtractWeeks subtracts the specified number of weeks from the given date.
func SubtractWeeks(date time.Time, weeks int) time.Time {
    return date.AddDate(0, 0, -weeks*7) 
}

// AddMonths adds the specified number of months to the given date.
func AddMonths(date time.Time, months int) time.Time {
    return date.AddDate(0, months, 0)
}

// SubtractMonths subtracts the specified number of months from the given date.
func SubtractMonths(date time.Time, months int) time.Time {
    return date.AddDate(0, -months, 0)
}

// AddYears adds the specified number of years to the given date.
func AddYears(date time.Time, years int) time.Time {
    return date.AddDate(years, 0, 0)
}

// SubtractYears subtracts the specified number of years from the given date.
func SubtractYears(date time.Time, years int) time.Time {
    return date.AddDate(-years, 0, 0)
}

// SumDates calculates the sum of two dates and returns the resulting date.
func SumDates(date1 time.Time, date2 time.Time) time.Time {
    return date1.Add(date2.Sub(date1))
}

// SubtractDates calculates the difference between two dates and returns the result as a time.Duration.
func SubtractDates(date1 time.Time, date2 time.Time) time.Duration {
    return date1.Sub(date2)
}

// GetAge calculates the age based on the given birthdate.
func GetAge(birthDate time.Time) int {
    current := time.Now()
    age := current.Year() - birthDate.Year()
    if current.YearDay() < birthDate.YearDay() {
        age--
    }
    return age
}
