package datetime

import "log"

// TimeAgoStr calculates the time difference from a given time string to the current time and returns a human-readable string.
func TimeAgoStr(fromTimeStr string) string {   
	fromTime, err := StringToDateTime(fromTimeStr)
	if err != nil {
		log.Printf("Error parsing fromTimeStr: %v", err)
		return "Invalid date"
	}
	return TimeAgo(fromTime)                 
}

// TimeAgoBetweenDatesStr calculates the time difference between two time strings and returns a string representing the duration.
func TimeAgoBetweenStr(startStr string, endStr string) string {
	startTime, err := StringToDateTime(startStr)
    if err != nil {
        log.Printf("Error parsing startStr: %v", err)
        return "Invalid start date"
    }
    endTime, err := StringToDateTime(endStr)
    if err != nil {
        log.Printf("Error parsing endStr: %v", err)
        return "Invalid end date"
    }
	return TimeAgoBetween(startTime, endTime) 
}
