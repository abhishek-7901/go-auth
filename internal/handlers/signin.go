package handlers

import (
	"go-auth/internal/database"
	"go-auth/internal/models"
	"time"

	"go-auth/pkg/generateTokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SigninRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SigninResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func HandleSignin(c *gin.Context) {
	var req SigninRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Find user by email
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := generateTokens.GenerateJWTtoken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate refresh token
	refreshToken, err := generateTokens.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}
	// Store refresh token in DB
	expiresAt := time.Now().Add(7 * 24 * time.Hour) // 7 days
	rt := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	}
	if err := database.DB.Create(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
		return
	}

	c.JSON(http.StatusOK, SigninResponse{Token: token, RefreshToken: refreshToken})

}
