package handlers

import (
	"github.com/gin-gonic/gin"
)

func ProtectedEndpoint(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	c.JSON(200, gin.H{
		"message": "You have accessed a protected route!",
		"user_id": userID,
		"email":   email,
	})
}
