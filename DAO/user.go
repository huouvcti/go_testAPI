package DAO

import (
	// "testAPI/models"
	"fmt"
	"testAPI/models"
)

type UserDAO struct {
}

func (dao UserDAO) CreateUser(user *models.User) error {
	fmt.Println("CreateUser")
	return DB.Create(user).Error
}

func (dao UserDAO) FindByID(id string) bool {
	fmt.Println("findID")

	var user models.User
	result := DB.Where("ID = ?", id).First(&user)

	if result.Error != nil {
		return false
	}

	fmt.Println(user)
	return true
}

func (dao UserDAO) LoginUser(id string, pw string) (bool, string, string) {
	fmt.Println("LoginUser")

	var user models.User
	result := DB.Where("ID = ? AND PW = ?", id, pw).First(&user)

	if result.Error != nil {
		return true, "", ""
	}

	fmt.Println(user)
	return false, user.ID, user.Name
}
