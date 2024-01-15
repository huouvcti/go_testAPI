package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type userToken struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type JwtUtil struct{}

func GetSecretKey() []byte {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := []byte(os.Getenv("SECRET_KEY"))

	return secretKey
}

func (jwtUtil *JwtUtil) GenerateToken(id string, name string) (string, error) {

	fmt.Println("scKey: ", GetSecretKey())
	claims := userToken{
		ID:   id,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // 토큰 만료 시간 설정 (1시간 후)
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetSecretKey())
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
		return GetSecretKey(), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
