package main

import (
	"fmt"
	"net/http"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// "fmt"
// "main/initializers"
// "net/http"

// "github.com/gin-gonic/gin"

//	func init (){
//		initializers.ConnectDb()
//	}
func main() {
	// fmt.Println("We are Live!!!")

	r := gin.Default();
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"mesage": "We are livey"})
	})
	err := r.Run(":8000");
	fmt.Println("🚀 Server running on port:", 8000)
	if err != nil {
	fmt.Println("Error running Server 👎")
	}
	fmt.Println("Server running speedily🍀")

	// router := gin.Default()

	//   router.GET("/", func(c *gin.Context) {
	//     c.JSON(http.StatusOK, gin.H{
	//       "message": "pongy",
	//     })
	//   })

	//   router.Run()

}
