package messages

func Error(message string) {
	Colorize(COLOR_RED, message)
}

func Info(message string) {
	Colorize(COLOR_CYAN, message)
}

func Success(message string) {
	Colorize(COLOR_GREEN, message)
}

func Warning(message string)  {
	Colorize(COLOR_YELLOW, message)
}

