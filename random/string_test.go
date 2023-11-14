package random

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(String(10))
	fmt.Println(String(100))
}