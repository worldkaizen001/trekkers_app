package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func InitiaRoute() *gin.Engine{
	r := gin.Default()
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/register",  controllers.Register)
	r.GET("users/:id", controllers.GetUserById)
	r.POST("users/search", controllers.SearchUserByName)
	r.DELETE("users/:id", controllers.DeletedUserById)
	r.PUT("users/:id", controllers.UpdateUser)
	r.GET("/test", controllers.Test)

	

	


	return r
}