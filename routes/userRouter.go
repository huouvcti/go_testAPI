package routes

import (
	"github.com/gin-gonic/gin"

	"testAPI/controllers"
)

func UserRouter(r *gin.Engine) {

	userController := controllers.User{}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/", userController.GetUser)

		userGroup.GET("/aaa", userController.GetAAA)
	}
}
