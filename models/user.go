package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt을 포함)
	ID         string `gorm:"size:20;unique;not null"` // 아이디
	PW         string `gorm:"size:20;not null"`        // 패스워드
	Name       string `gorm:"size:10;not null"`        // 이름
	IsActive   bool   `gorm:"default:true"`            // 활성 상태
}
