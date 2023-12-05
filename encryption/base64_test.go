package encryption

import (
	"fmt"
	"testing"
)


const (
	originalStrBase64 = "This is the original String"
	encodedStrBase64 = "VGhpcyBpcyB0aGUgb3JpZ2luYWwgU3RyaW5n"
)


func TestEncodeBase64(t *testing.T) {
	got := encodedStrBase64
	expected := EncodeBase64(originalStrBase64) 
	fmt.Println(got, expected)
	if got != expected {
		t.Errorf("Encoding failed '%s' != %s", got, expected)
	}
}


func TestDecodeBase64(t *testing.T) {
	got := originalStrBase64
	expected := DecodeBase64(encodedStrBase64)
	if got != expected {
		t.Errorf("Encoding failed '%s' != %s", got, expected)
	}
}