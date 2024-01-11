package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	GetUser(c *gin.Context)

	Join(c *gin.Context)
}

type UserController struct {
}

func (user *UserController) GetUser(c *gin.Context) {
	data := make(map[string]interface{})

	data["name"] = "kim"
	data["age"] = 25

	c.JSON(200, data)
}

func (user *UserController) Join(c *gin.Context) {
	body := c.Request.Body

	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var data map[string]interface{}
	json.Unmarshal(value, &data)

	fmt.Println(value)

	c.JSON(200, data)
}
