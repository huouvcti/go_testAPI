package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"testAPI/config"
	"testAPI/routes"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB 연결
	db := config.ConnDB()
	config.SetORM(db)

	r := gin.Default()

	// API Routing
	apiGroup := r.Group("/api")
	{
		routes.V1Group(apiGroup)
	}

	// 테스트 페이지
	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/join", func(c *gin.Context) {
		c.HTML(http.StatusOK, "join.html", gin.H{})
	})

	runPort := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%s", runPort)) // 포트 설정
}
