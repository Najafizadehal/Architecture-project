package main

import (
	"architecture/initializers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	// r := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
		
}
