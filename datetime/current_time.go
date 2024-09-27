package datetime

import "time"

// CurrentDate returns the current date as a formatted string in the format "2006-01-02".
func CurrentDate() string {
	return time.Now().Format(dateLayout)
}

// CurrentDateTime returns the current date and time as a formatted string in the format "2006-01-02 15:04:05".
func CurrentDateTime() string {
	return time.Now().Format(dateTimeLayout)
}

// CurrentYear returns the current year as a formatted string.
func CurrentYear() string {
	return time.Now().Format(yearLayout)
}

// CurrentTime returns the current time as a formatted string in the format "15:04:05".
func CurrentTime() string {
	return time.Now().Format(secondsLayout)
}

// CurrentHour returns the current hour as a formatted string in the format "15".
func CurrentHour() string {
	return time.Now().Format(hourLayout)
}
