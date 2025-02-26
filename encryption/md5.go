package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

// EncodeMD5 encodes the given text using MD5 hashing algorithm.
func EncodeMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// DecodeMD5 decodes the given hexadecimal encoded MD5 hash.
func DecodeMD5(encoded string) string {
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(decoded)
}
