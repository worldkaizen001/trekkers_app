package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword encrypts a password
func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashed), err
}

// CheckPassword compares a hashed password with a plaintext one
func CheckPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
