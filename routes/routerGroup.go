package routes

import (
	"github.com/gin-gonic/gin"
)

func V1Group(r *gin.RouterGroup) {

	v1Group := r.Group("/v1")
	{
		UserRouter(v1Group)
	}
}
