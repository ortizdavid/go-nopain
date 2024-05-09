package messages

import (
	"testing"
)

type colorTest struct {
	color string
	text string
}


func Test_PrintColorizeText(t *testing.T) {
	testCases := []colorTest {
		{ color: COLOR_BLACK, text: "Black Text" },
		{ color: COLOR_WHITE, text: "White Text" },
		{ color: COLOR_CYAN, text: "Cyan Text" },
		{ color: COLOR_YELLOW, text: "Yellow Text" },
		{ color: COLOR_RED, text: "Red Text" },
		{ color: COLOR_GREEN, text: "Green Text" },
		{ color: COLOR_MAGENTA, text: "Mangenta Text" },
		{ color: COLOR_BRIGHT_BLACK, text: "Bright Black Text" },
		{ color: COLOR_BRIGHT_WHITE, text: "Bright White Text" },
		{ color: COLOR_BRIGHT_CYAN, text: "Bright Cyan Text" },
		{ color: COLOR_BRIGHT_YELLOW, text: "Bright Yellow Text" },
		{ color: COLOR_BRIGHT_RED, text: "Bright Red Text" },
		{ color: COLOR_GREEN, text: "Bright Green Text" },
		{ color: COLOR_BRIGHT_MAGENTA, text: "Bright Mangenta Text" },
	}
	for _, test := range testCases {
		Colorize(test.color, test.text)
	}
}