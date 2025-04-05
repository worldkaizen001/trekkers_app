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
	// "main/initializers"
)

func CreateUser (c *gin.Context) {
    
	type Body struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`

	}
	var users Body;

	err := c.ShouldBind(&users);
	if err != nil {
		c.JSON(400, gin.H{"error": "err"})
	}

	newUser:= models.User{Name: users.Name, Email: users.Email, Password: users.Password}
	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		fmt.Println("error occured inserting into DB")
		return
	}
	c.JSON(200, gin.H{"success": "user created successfully"})
	

	
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
	

    if err := initializers.DB.Create(&user).Error; 
	err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User could not be created"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

type UserResponseDTO struct {
	ID      uint            `json:"id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	// Profile models.Profile  `json:"profile"`
	// Steps   []models.Step   `json:"steps"`
	// Password string `json:"password"`
}
func GetAllUsers(c *gin.Context){
	var user []models.User
err :=	initializers.DB.Find(&user).Error;
if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting users"});
	return
}

c.JSON(http.StatusOK, gin.H{
	"success": true,
	"data": user,
})
}

func GetUserById(c *gin.Context){
	id:= c.Param("id")
	var user models.User

err := initializers.DB.First(&user, id).Error;
if err != nil {
  c.JSON(http.StatusBadRequest, gin.H{"Error": err})
  return
}
res := UserResponseDTO{
	ID:      user.ID,
	Name:    user.Name,
	Email:   user.Email,
	// Password: user.Password,
	// Password: user.Password,
	// Profile: user.Profile,
	// Steps:   user.Steps,
}
c.JSON(http.StatusOK, gin.H{
	"status": true,
	"data": res,
})
}
func DeletedUserById(c *gin.Context){
	id:= c.Param("id")
	var user models.User;
 err :=	initializers.DB.Delete(&user, id).Error;
 if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "error occured",
	})
 }

 c.JSON(200, gin.H{"success": "User deleted successfully"})
}
func SearchUserByName(c *gin.Context){
	name :=  c.Query("name")
	var user models.User;
	err :=	initializers.DB.First(&user, "name = ?", name).Error;
 if err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "error occured",
	})
 }

 c.JSON(200, gin.H{
	"success": true,
	"data": user,
})

}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Find the user first
	if err := initializers.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind the JSON input
	var input struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update the email field
	if err := initializers.DB.Model(&user).Update("Email", input.Email).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email updated successfully", "user": user})
}






