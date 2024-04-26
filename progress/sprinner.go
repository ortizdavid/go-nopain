package progress

import (
	"fmt"
	"time"
)

// Spinner displays a spinning animation to indicate progress.
// It takes a time duration delay as input to control the speed of the animation.
func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r) // Print the spinner character
			time.Sleep(delay)      // Pause for the specified duration
		}
	}
}