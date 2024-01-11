package config

import (
	"log"
	"testAPI/models"

	"gorm.io/gorm"
)

func SetORM(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("모델 마이그레이션 실패: %v", err)
	}
}
