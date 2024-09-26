package datetime

// TimeAgoStr calculates the time difference from a given time string to the current time and returns a human-readable string.
func TimeAgoStr(fromTimeStr string) string {     
	return TimeAgo(StringToDateTime(fromTimeStr))                 
}

// TimeAgoBetweenDatesStr calculates the time difference between two time strings and returns a string representing the duration.
func TimeAgoBetweenDatesStr(startStr string, endStr string) string {
	return TimeAgoBetweenDates(StringToDateTime(startStr), StringToDateTime(endStr)) 
}
