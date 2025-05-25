package generateTokens

import (
	"go-auth/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("your-secret-key") // This should come from config in production

// GenerateToken generates a JWT token for a user
func GenerateJWTtoken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
