package encryption

import (
	"encoding/base64"
	"log"
)


func EncodeBase64(text string) string {
	byteText := []byte(text)
	encoded := base64.StdEncoding.EncodeToString(byteText)
	return encoded
}


func DecodeBase64(strEncoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(strEncoded)
	if err != nil {
		log.Fatal(err.Error())
	}
	text := string(decoded)
	return text
}


func CompareBase64Encoding(encodedStr string, originalStr string) bool {
	decodedStr := DecodeBase64(encodedStr) 
	return encodedStr == decodedStr
}