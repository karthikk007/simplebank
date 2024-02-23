package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the brypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

// ValidatePassword validates password with hashed password
func ValidatePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
