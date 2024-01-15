package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type userToken struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type JwtUtil struct{}

func (jwtUtil *JwtUtil) GenerateToken(id string, name string) (string, error) {
	claims := userToken{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // 토큰 만료 시간 설정 (1시간 후)
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtUtil *JwtUtil) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("올바르지 않은 토큰 알고리즘")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
