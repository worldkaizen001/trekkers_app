package controllers

import (
	// "fmt"
	// "main/models"

	"fmt"
	"main/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "main/initializers"
)
var DBB *gorm.DB
func CreateUser (c *gin.Context) {
    
	type Body struct {
		Name string `json:"name"`
		Email string `json:"email"`

	}
	var users Body;

	err := c.ShouldBind(&users);
	if err != nil {
		c.JSON(400, gin.H{"error": "err"})
	}

	newUser:= models.User{Name: users.Name, Email: users.Email}
	result := DBB.Create(&newUser)
	if result.Error != nil {
		fmt.Println("error occured inserting into DB")
		return
	}
	c.JSON(200, gin.H{"success": "user created "})
	

	
}

