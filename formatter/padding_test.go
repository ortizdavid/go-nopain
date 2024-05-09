package formatter

import "testing"


func Test_LeftPadding(t *testing.T) {
	original := "123"
	allowedLength := 8
	char := "0"
	expected := "00000123"
	result := LeftPadWithChar(original, allowedLength, char)
	if result != expected {
		t.Errorf("\nLeftPadWitchar(%s, %d, %s) = %s. got: %s", original, allowedLength, char, result, expected)
	}
}


func Test_RightPadding(t *testing.T) {
	original := "test"
	allowedLength := 12
	char := "*"
	expected := "test********"
	result := RightPadWithChar(original, allowedLength, char)
	if result != expected {
		t.Errorf("\nLeftPadWitchar(%s, %d, %s) = %s. got: %s", original, allowedLength, char, result, expected)
	}
}