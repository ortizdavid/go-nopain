package encryption

import (
	"encoding/base64"
	"log"
)


func StringToBase64(text string) string {
	byteText := []byte(text)
	encoded := base64.StdEncoding.EncodeToString(byteText)
	return encoded
}


func Base64ToString(strEncoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(strEncoded)
	if err != nil {
		log.Fatal(err.Error())
	}
	text := string(decoded)
	return text
}


func CompareBase64Encoding(encodedStr string, originalStr string) bool {
	decodedStr := Base64ToString(encodedStr) 
	return encodedStr == decodedStr
}