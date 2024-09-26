package datetime


// TimeAgoStr calculates the time difference from a given time string to the current time and returns a human-readable string.
func TimeAgoStr(fromTimeStr string) string {     
	fromTime := StringToDateTime(fromTimeStr) // Convert the time string to a time.Time object.
	return TimeAgo(fromTime)                    // Use the TimeAgo function to get the difference.
}

// TimeAgoBetweenDatesStr calculates the time difference between two time strings and returns a string representing the duration.
func TimeAgoBetweenDatesStr(startStr string, endStr string) string {
	return TimeAgoBetweenDates(StringToDateTime(startStr), StringToDateTime(endStr)) // Convert strings to time.Time and calculate the difference.
}
