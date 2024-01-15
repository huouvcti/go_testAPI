package routes

import (
	"testAPI/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	var userController controllers.UserInterface = &controllers.UserController{}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/", userController.GetUser)

		userGroup.POST("/join", userController.Join)

		userGroup.POST("/login", userController.Login)
	}
}
