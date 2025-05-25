package main

import (
	"fmt"
	"go-auth/internal/database"
	"go-auth/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	database.InitDB()
	router := gin.Default()

	router.POST("api/auth/signup", handlers.HandleSignup)
	router.POST("api/auth/signin", handlers.HandleSignin)
	// router.POST("api/auth/refresh", handlers.HandleSignUp)
	// router.POST("api/auth/revoke", handlers.HandleSignUp)
	// router.POST("api/auth/signup", handlers.HandleSignUp)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server", err)
	}
}
