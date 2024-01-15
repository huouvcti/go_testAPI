package middleware

import (
	"testAPI/utils"

	"github.com/gin-gonic/gin"
)

var jwtUtil = utils.JwtUtil{}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "인증이 필요합니다."})
			c.Abort()
			return
		}

		// 토큰 검증
		validatedToken, err := jwtUtil.VerifyToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "토큰 검증 실패"})
			c.Abort()
			return
		}

		// 검증된 토큰의 클레임을 가져와서 사용할 수 있음
		c.Set("token_claims", validatedToken.Claims)
		c.Next()
	}
}
