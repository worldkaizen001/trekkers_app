package main

import (
	"fmt"
	"main/controllers"
	"main/initializers"
	"net/http"

	// "net/http"

	"github.com/gin-gonic/gin"
)

// "fmt"
// "main/initializers"
// "net/http"

// "github.com/gin-gonic/gin"

func init() {
	initializers.ConnectDb()
}
func main() {
	// fmt.Println("We are Live!!!")

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"mesage": "We are live"})
	})
	err := r.Run(":8000")
	fmt.Println("🚀 Server running on port:", 8000)
	if err != nil {
		fmt.Println("Error running Server 👎")
	}
	fmt.Println("Server running speedily🍀")

	r.POST("/create", controllers.CreateUser)

	

}
