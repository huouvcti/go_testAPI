package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt을 포함)
	ID         string `gorm:"size:20;unique;not null"` // 사용자 이름
	PW         string `gorm:"size:20;not null"`        // 이메일, 고유 값
	Name       string `gorm:"size:10;not null"`        // 나이, 기본값은 18
	IsActive   bool   `gorm:"default:true"`            // 활성 상태
}
