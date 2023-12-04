package encryption

import (
	"fmt"
	"testing"
)

func TestStringToMd5(t *testing.T) {
	text := "Hello World"
	fmt.Println(StringToMd5(text))
}

func TestMd5ToString(t *testing.T) {
	encoded := "b10a8db164e0754105b7a99be72e3fe5"
	fmt.Println(Md5ToString(encoded))
}