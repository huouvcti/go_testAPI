package DAO

import (
	// "testAPI/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDAO(db *gorm.DB) {
	DB = db
}
