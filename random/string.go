package random

import (
	"math/rand"
)

// charset defines the characters used for generating random strings.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String generates a random string of the specified length using characters from the charset.
func String(length int) string {
	seedRandom()
	if length <= 0 {
		panic("Invalid string length. Length should be greater than zero.")
	}
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
