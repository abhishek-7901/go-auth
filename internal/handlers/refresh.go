package handlers

import (
	"go-auth/internal/database"
	"go-auth/internal/models"
	"go-auth/pkg/generateTokens"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func HandleRefresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	var rt models.RefreshToken
	if err := database.DB.Where("token = ? AND revoked = ? AND expires_at > ?", req.RefreshToken, false, time.Now()).First(&rt).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}
	var user models.User
	if err := database.DB.First(&user, rt.UserID).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	// Generate new access token
	token, err := generateTokens.GenerateJWTtoken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Optionally, rotate refresh token (best practice)
	newRefreshToken, err := generateTokens.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}
	rt.Token = newRefreshToken
	rt.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)
	if err := database.DB.Save(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update refresh token"})
		return
	}
	c.JSON(http.StatusOK, RefreshResponse{Token: token, RefreshToken: newRefreshToken})
}
