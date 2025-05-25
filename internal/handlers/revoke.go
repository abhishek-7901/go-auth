package handlers

import (
	"go-auth/internal/database"
	"go-auth/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RevokeRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func HandleRevoke(c *gin.Context) {
	var req RevokeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	var rt models.RefreshToken
	if err := database.DB.Where("token = ? AND revoked = ?", req.RefreshToken, false).First(&rt).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	rt.Revoked = true
	if err := database.DB.Save(&rt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Refresh token revoked"})
}
