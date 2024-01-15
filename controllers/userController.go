package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"testAPI/models"
	"testAPI/utils"

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
var jwtUtil = utils.JwtUtil{}

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

	err, id, name := userDAO.LoginUser(loginReq.ID, loginReq.PW)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "아이디 또는 패스워드가 일치하지 않음"})
		return
	}

	/* JWT 사용 X

	fmt.Println(id, name)

	type Data struct {
		ID   string
		Name string
	}
	d := Data{id, name}

	c.JSON(200, gin.H{"msg": "로그인 성공", "data": d})

	*/

	token, err := jwtUtil.GenerateToken(id, name)
	if err != nil {
		c.JSON(500, gin.H{"error": "토큰 생성 실패"})
		return
	}

	c.JSON(200, gin.H{"msg": "로그인 성공", "token": token})
}
