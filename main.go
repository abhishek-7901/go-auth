package main

import (
	"/go-auth/internal/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	router := gin.Default()

	router.post("api/auth/signup", handlers.HandleSignUp)
}
