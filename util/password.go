package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash password: %w", err)
	}
	return string(HashedPassword), nil
}

// CheckPassword checks the password
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}