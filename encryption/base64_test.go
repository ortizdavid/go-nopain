package encryption

import (
	"fmt"
	"testing"
)

const (
	originalStr = "This is the original String"
	encodedStr = "VGhpcyBpcyB0aGUgb3JpZ2luYWwgU3RyaW5n"
)


func TestStringToBase64(t *testing.T) {
	got := encodedStr
	expected := StringToBase64(originalStr) 
	fmt.Println(got, expected)
	if got != expected {
		t.Errorf("Encoding failed '%s' != %s", got, expected)
	}
}

func TestBase64ToString(t *testing.T) {
	got := originalStr
	expected := Base64ToString(encodedStr)
	if got != expected {
		t.Errorf("Encoding failed '%s' != %s", got, expected)
	}
}