package handlers

import (
	"go-auth/internal/database"
	"go-auth/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignupResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func HandleSignup(c *gin.Context) {
	var req SignupRequest

	// Parse and validate the request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Check if user already exists
	var existingUser models.User
	result := database.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	} else if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error checking for existing user",
		})
		return
	}

	// Create a new user
	user := models.User{
		Email:    req.Email,
		Password: req.Password, // Will be hashed by BeforeCreate hook
	}

	// Save to database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Create response
	response := SignupResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	// Send response
	c.JSON(http.StatusCreated, response)
}
