package main

import (
	"github.com/gin-gonic/gin"

	"testAPI/routes"
)

func main() {
	r := gin.Default()

	routes.UserRouter(r)

	r.Run(":1234") // 포트 설정
}
