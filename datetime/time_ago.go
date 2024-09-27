package datetime

import (
	"fmt"
	"time"
)

// TimeAgo calculates the time difference from a given time to the current time and returns a human-readable string.
func TimeAgo(fromTime time.Time) string {
	now := time.Now()             
	diff := now.Sub(fromTime) 

	// Calculate the time difference in different units.
	years := int(diff.Hours() / 24 / 365)
	months := int(diff.Hours() / 24 / 30) 
	days := int(diff.Hours() / 24)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes())
	seconds := int(diff.Seconds())

	// Return the appropriate time difference string.
	if years >= 1 {
		return fmt.Sprintf("%d years ago", years)
	} else if months >= 1 {
		return fmt.Sprintf("%d months ago", months)
	} else if days >= 1 {
		return fmt.Sprintf("%d days ago", days)
	} else if hours >= 1 {
		return fmt.Sprintf("%d hours ago", hours)
	} else if minutes >= 1 {
		return fmt.Sprintf("%d minutes ago", minutes)
	} else {
		return fmt.Sprintf("%d seconds ago", seconds)
	}
}

// TimeAgoBetweenDates calculates the time difference between two time values and returns a string representing the duration.
func TimeAgoBetween(start time.Time, end time.Time) string {
	diff := end.Sub(start)                    
	return TimeAgo(start) + " to " + TimeAgo(end) + " (" + diff.String() + ")" 
}
