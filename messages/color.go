package messages

import (
	"fmt"
)

// ANSI color escape sequences for text colorization.
const (
	COLOR_RESET             = "\u001b[0m"
	COLOR_BLACK             = "\u001b[30m"
	COLOR_RED               = "\u001b[31m"
	COLOR_GREEN             = "\u001b[32m"
	COLOR_YELLOW            = "\u001b[33m"
	COLOR_BLUE              = "\u001b[34m"
	COLOR_MAGENTA           = "\u001b[35m"         // Corrected typo in color name
	COLOR_CYAN              = "\u001b[36m"
	COLOR_WHITE             = "\u001b[37m"
	COLOR_BRIGHT_BLACK      = "\u001b[30;1m"
	COLOR_BRIGHT_RED        = "\u001b[31;1m"
	COLOR_BRIGHT_GREEN      = "\u001b[32;1m"
	COLOR_BRIGHT_YELLOW     = "\u001b[33;1m"
	COLOR_BRIGHT_BLUE       = "\u001b[34;1m"
	COLOR_BRIGHT_MAGENTA    = "\u001b[35;1m"     // Corrected typo in color name
	COLOR_BRIGHT_CYAN       = "\u001b[36;1m"
	COLOR_BRIGHT_WHITE      = "\u001b[37;1m"
)

// Colorize prints the given text in the specified color.
func Colorize(color string, text string) {
	fmt.Println(string(color), text, string(COLOR_RESET))
}

// Colorizef prints the formatted text in the specified color.
func Colorizef(color string, formattedText string, values ...interface{}) {
	fmt.Printf(string(color), formattedText, values, string(COLOR_RESET))
}
