package datetime

import "time"

// CurrentDate returns the current date as a formatted string in the format "2006-01-02".
func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// CurrentDateTime returns the current date and time as a formatted string in the format "2006-01-02 15:04:05".
func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// CurrentYear returns the current year as a formatted string.
func CurrentYear() string {
	return time.Now().Format("2006")
}

// CurrentTime returns the current time as a formatted string in the format "15:04:05".
func CurrentTime() string {
	return time.Now().Format("15:04:05")
}

// CurrentHour returns the current hour as a formatted string in the format "15".
func CurrentHour() string {
	return time.Now().Format("15")
}
