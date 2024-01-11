package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type DBORM struct {
	*gorm.DB
}

func ConnDB() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 사용자명:비밀번호@tcp(호스트:포트)/데이터베이스명?옵션
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	// 연결 테스트
	var result int
	db.Raw("SELECT 1").Scan(&result)

	if result == 1 {
		log.Println("데이터베이스 연결 성공")
	}

	return db
}
