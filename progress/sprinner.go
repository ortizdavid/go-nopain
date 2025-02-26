package progress

import (
	"fmt"
	"time"
)

// Spinner displays a spinning animation to indicate progress.
func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // Print the spinner character
			time.Sleep(delay)      
		}
	}
}