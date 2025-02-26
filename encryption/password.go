package encryption

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)


// HashPassword generates a bcrypt hash from the given password.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(hash)
}

// CheckPassword compares a bcrypt hashed password with a plain-text password.
func CheckPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}