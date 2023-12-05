package encryption

import (
	"fmt"
	"testing"
)

const (
	originalStrMd5 = "Hello World"
	encodedStrMd5 = "b10a8db164e0754105b7a99be72e3fe5"
)

func TestEncodeMD5(t *testing.T) {
	text := originalStrMd5
	fmt.Println(EncodeMD5(text))
}


func TestDecodeMD5(t *testing.T) {
	encoded := encodedStrMd5
	fmt.Println(DecodeMD5(encoded))
}