package token

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRefreshToken generates a secure random string for refresh tokens
func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
