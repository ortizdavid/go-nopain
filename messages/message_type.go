package messages

// Error prints the given message in red color to indicate an error.
func Error(message string) {
	Colorize(COLOR_RED, message)
}

// Info prints the given message in cyan color to provide informational messages.
func Info(message string) {
	Colorize(COLOR_CYAN, message)
}

// Success prints the given message in green color to indicate successful operations.
func Success(message string) {
	Colorize(COLOR_GREEN, message)
}

// Warning prints the given message in yellow color to indicate warnings.
func Warning(message string) {
	Colorize(COLOR_YELLOW, message)
}