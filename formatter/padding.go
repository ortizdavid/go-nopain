package formatter

import "strings"

// Add repeated chars to the left of a string value
// Returns the concatenation of Padding and Original value
// Padding: char reptead N times
func LeftPadWithChar(originalValue string, allowedLenght int, char string) string {
	originalLength := len(originalValue)

	if allowedLenght <= 0 || char == ""  || originalLength == 0 {
		return ""
	}

	if  originalLength >= allowedLenght {
		return originalValue
	} 

	paddingLength := allowedLenght - originalLength
	padding := strings.Repeat(char, paddingLength)
	return padding + originalValue
}


// Add repeated chars to the left of a string value
// Returns the concatenation of Padding and Original value
// Padding: char reptead N times
func RightPadWithChar(originalValue string, allowedLenght int, char string) string {
	originalLength := len(originalValue)

	if allowedLenght <= 0 || char == ""  || originalLength == 0 {
		return ""
	}

	if  originalLength >= allowedLenght {
		return originalValue
	} 

	paddingLength := allowedLenght - originalLength
	padding := strings.Repeat(char, paddingLength)
	return originalValue + padding
}
