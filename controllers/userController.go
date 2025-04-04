package controllers

import (
	// "fmt"
	// "main/models"

	"fmt"
	"main/initializers"
	"main/models"
	"main/utils"
	"net/http"

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
	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		fmt.Println("error occured inserting into DB")
		return
	}
	c.JSON(200, gin.H{"success": "user created "})
	

	
}

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password before storing
    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword

    if err := initializers.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User could not be created"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}