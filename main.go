package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Create a route group
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health-check", healthCheck)
	}

	// Start the server on port 8080
	router.Run(":9000")
}

// healthCheck is the handler function for the /health-check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"message": "Health check successful",
	})
}
