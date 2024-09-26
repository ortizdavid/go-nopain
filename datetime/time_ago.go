package datetime

import (
	"fmt"
	"time"
)

// TimeAgo calculates the time difference from a given time to the current time and returns a human-readable string.
func TimeAgo(fromTime time.Time) string {
	now := time.Now()                       // Get the current time.
	diff := now.Sub(fromTime)               // Calculate the difference between now and the provided time.

	// Calculate the time difference in various units.
	years := diff.Hours() / 24 / 365
	months := years * 12
	days := diff.Hours() / 24
	hours := diff.Hours()
	minutes := diff.Minutes()
	seconds := diff.Seconds()

	// Return the appropriate time difference string.
	if years >= 1 {
		return fmt.Sprintf("%.0f years ago", years)
	} else if months >= 1 {
		return fmt.Sprintf("%.0f months ago", months)
	} else if days >= 1 {
		return fmt.Sprintf("%.0f days ago", days)
	} else if hours >= 1 {
		return fmt.Sprintf("%.0f hours ago", hours)
	} else if minutes >= 1 {
		return fmt.Sprintf("%.0f minutes ago", minutes)
	} else {
		return fmt.Sprintf("%.0f seconds ago", seconds)
	}
}

// TimeAgoBetweenDates calculates the time difference between two time values and returns a string representing the duration.
func TimeAgoBetween(start time.Time, end time.Time) string {
	diff := end.Sub(start)                     // Calculate the difference between the start and end times.
	return TimeAgo(start) + " to " + TimeAgo(end) + " (" + diff.String() + ")" 
}
