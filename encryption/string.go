package encryption

import (
	"log"
	"time"
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
)


func GenerateUUID() string {
	uniqueId := uuid.New()
	return uniqueId.String()
}


func GenerateRandomToken() string {
	length := 100
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err.Error())
	}
	token := base64.RawURLEncoding.EncodeToString(randomBytes)
	return token
}


func GenerateCode(prefix string) string {
	timestamp := time.Now().Format("20060102150405") 
	return prefix + timestamp  
}
