package controllers

import (
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (user User) GetUser(c *gin.Context) {
	data := make(map[string]interface{})

	data["name"] = "kim"
	data["age"] = 25

	c.JSON(200, data)
}

func (user User) GetAAA(c *gin.Context) {
	data := make(map[string]interface{})

	data["name"] = "AAA"
	data["age"] = 0

	c.JSON(200, data)
}
