package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"testAPI/models"

	"testAPI/DAO"
)

type UserInterface interface {
	GetUser(c *gin.Context)

	Join(c *gin.Context)
	Login(c *gin.Context)
}

type UserController struct {
}

var userDAO = DAO.UserDAO{}

func (user *UserController) GetUser(c *gin.Context) {
	data := make(map[string]interface{})

	data["name"] = "kim"
	data["age"] = 25

	c.JSON(200, data)
}

func (user *UserController) Join(c *gin.Context) {

	var joinReq models.User

	if err := c.BindJSON(&joinReq); err != nil {
		log.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if joinReq.ID == "" || joinReq.PW == "" || joinReq.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "not enough"})
		return
	}

	fmt.Println(joinReq.ID, joinReq.PW, joinReq.Name)

	if userDAO.FindByID(joinReq.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "아이디가 이미 존재함"})
		return
	}

	userDAO.CreateUser(&joinReq)

	c.JSON(200, gin.H{"msg": "회원가입 성공"})
}

func (user *UserController) Login(c *gin.Context) {

	var loginReq models.User

	if err := c.BindJSON(&loginReq); err != nil {
		log.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if loginReq.ID == "" || loginReq.PW == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "not enough"})
		return
	}

	err, id, pw := userDAO.LoginUser(loginReq.ID, loginReq.PW)
	if err {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "아이디 또는 패스워드가 일치하지 않음"})
		return
	}

	fmt.Println(id, pw)

	type Data struct {
		ID   string
		Name string
	}

	d := Data{id, pw}

	c.JSON(200, gin.H{"msg": "로그인 성공", "data": d})
}
