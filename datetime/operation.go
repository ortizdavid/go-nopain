package datetime 

import "time"


func AddDaysToDate(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, days)
}

func SubtractDaysFromDate(date time.Time, days int) time.Time {
	return date.AddDate(0, 0, -days)
}

func AddToTime(date time.Time, years, months, days int) time.Time {
	return date.AddDate(years, months, days)
}

func SumDates(date1, date2 time.Time) time.Time {
	return date1.Add(date2.Sub(date1))
}

func GetAge(birthdate time.Time) int {
    current := time.Now()
    age := current.Year() - birthdate.Year()

    // Check if the birthday for this year has already occurred.
    if current.YearDay() < birthdate.YearDay() {
        age--
    }
    return age
}