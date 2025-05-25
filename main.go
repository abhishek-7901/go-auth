package main

import (
	"fmt"
	"go-auth/internal/database"
	"go-auth/internal/handlers"
	"go-auth/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	database.InitDB()
	router := gin.Default()

	router.POST("api/auth/signup", handlers.HandleSignup)
	router.POST("api/auth/signin", handlers.HandleSignin)
	router.POST("api/auth/refresh", handlers.HandleRefresh)
	router.POST("api/auth/revoke", handlers.HandleRevoke)

	router.GET("/api/protected", middleware.AuthMiddleware(), handlers.ProtectedEndpoint)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server", err)
	}
}
