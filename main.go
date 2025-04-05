package main

import (
	"fmt"
	"main/initializers"
	"main/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectDb()
}

func main() {
	r := routes.InitiaRoute()
	// Register root route
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"mesage": "We are live40"})
	})

	// Setup other routes before starting the server

	fmt.Println("🚀 Server running on port: 8000")
	fmt.Println("Server running speedily🍀")

	// Start the server (blocking call)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error running Server 👎")
	}
}
