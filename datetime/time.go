package datetime

import (
	"fmt"
	"time"
)


func TimeAgo(from time.Time) string {
	now := time.Now()
	diff := now.Sub(from)

	years := diff.Hours() / 24 / 365
	months := years * 12
	days := diff.Hours() / 24
	hours := diff.Hours()
	minutes := diff.Minutes()
	seconds := diff.Seconds()

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


func TimeAgoBetweenDates(start time.Time, end time.Time) string {
	diff := end.Sub(start)
	return TimeAgo(start) + " to " + TimeAgo(end) + " (" + diff.String() + ")"
}
