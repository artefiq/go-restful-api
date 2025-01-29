package models

import "github.com/golang-jwt/jwt/v4"

// Claims structure for JWT
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
