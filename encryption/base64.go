package encryption

import (
	"encoding/base64"
	"log"
)

// EncodeBase64 encodes the given text to Base64.
// It takes a string input and returns the Base64 encoded string.
func EncodeBase64(text string) string {
	byteText := []byte(text)
	encoded := base64.StdEncoding.EncodeToString(byteText)
	return encoded
}

// DecodeBase64 decodes the given Base64 encoded string.
// It takes a Base64 encoded string as input and returns the decoded text.
func DecodeBase64(strEncoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(strEncoded)
	if err != nil {
		log.Fatal(err.Error())
	}
	text := string(decoded)
	return text
}

// CompareBase64Encoding compares a Base64 encoded string with its original text.
// It takes an encoded string and its original string as input,
// decodes the encoded string, and compares it with the original string.
func CompareBase64Encoding(encodedStr string, originalStr string) bool {
	decodedStr := DecodeBase64(encodedStr)
	return decodedStr == originalStr
}
