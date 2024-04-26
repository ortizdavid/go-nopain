package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

// EncodeMD5 encodes the given text using MD5 hashing algorithm.
// It takes a string input and returns the MD5 hash as a hexadecimal string.
func EncodeMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// DecodeMD5 decodes the given hexadecimal encoded MD5 hash.
// It takes a hexadecimal encoded MD5 hash as input and returns the decoded string.
func DecodeMD5(encoded string) string {
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(decoded)
}
