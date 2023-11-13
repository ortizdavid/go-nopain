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

func GenerateRandomToken() (string, error) {
	length := 100
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Println("Error generating random bytes:", err)
		return "", err
	}
	token := base64.RawURLEncoding.EncodeToString(randomBytes)
	return token, nil
}

func GenerateCode(prefix string) string {
	timestamp := time.Now().Format("20060102150405") 
    return prefix + timestamp  
}
