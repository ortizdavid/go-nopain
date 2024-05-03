package random

import (
	"testing"
)

func TestString(t *testing.T) {
	// Test case for positive length
	length := 10
	result := String(length)
	if len(result) != length {
		t.Errorf("String(%d) returned a string of length %d, expected %d", length, len(result), length)
	}
	// Test case for zero length
	length = 0
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("String(%d) did not panic as expected", length)
		}
	}()
	String(length)
	// Test case for negative length
	length = -5
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("String(%d) did not panic as expected", length)
		}
	}()
	String(length)
}
