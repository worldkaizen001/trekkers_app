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

	foods := [3]string{"rice","beans", "yam",}
	to := make([]int, 4)
	



	for i, v := range to {
     fmt.Print(i, v)
	}
	for iv := 0; iv < len(to); iv++{

	}

	for i, v := range foods {
		fmt.Println(i,v)
	}

	// Setup other routes before starting the server

	fmt.Println("🚀 Server running on port: 8000")
	fmt.Println("Server running speedily🍀")

	// Start the server (blocking call)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error running Server 👎")
	}
}
