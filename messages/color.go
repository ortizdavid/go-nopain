package messages

import (
	"fmt"
)

const (
	COLOR_RESET               = "\u001b[0m"
	COLOR_BLACK               = "\u001b[30m"
	COLOR_RED                 = "\u001b[31m"
	COLOR_GREEN               = "\u001b[32m"
	COLOR_YELLOW              = "\u001b[33m"
	COLOR_BLUE                = "\u001b[34m"
	COLOR_MANGENTA            = "\u001b[35m"
	COLOR_CYAN                = "\u001b[36m"
	COLOR_WHITE               = "\u001b[37m"
	COLOR_BRIGHT_BLACK        = "\u001b[30;1m"
	COLOR_BRIGHT_RED          = "\u001b[31;1m"
	COLOR_BRIGHT_GREEN        = "\u001b[32;1m"
	COLOR_BRIGHT_YELLOW       = "\u001b[33;1m"
	COLOR_BRIGHT_BLUE         = "\u001b[34;1m"
	COLOR_BRIGHT_MANGENTA     = "\u001b[35;1m"
	COLOR_BRIGHT_CYAN         = "\u001b[36;1m"
	COLOR_BRIGHT_WHITE        = "\u001b[37;1m"
	
)

func Colorize(color string, text string) {
	fmt.Println(string(color), text, string(COLOR_RESET))
}

func Colorizef(color string, formatedText string, values ...any) {
	fmt.Printf(string(color), formatedText, values, string(COLOR_RESET))
}
