package encryption

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)


// HashPassword generates a bcrypt hash from the given password.
// It takes a string password input and returns the bcrypt hash as a string.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(hash)
}

// CheckPassword compares a bcrypt hashed password with a plain-text password.
// It takes the hashed password and the plain-text password as inputs,
// returning true if the passwords match, and false otherwise.
func CheckPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}