package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashed), err
}

// CheckPassword checks if the provided password matches the hash
func CheckPassword(password, hashed string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
