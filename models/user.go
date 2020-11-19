package models

import (
	"errors"
	"fmt"
	"go-pangu/db"

	"gorm.io/gorm"
)

//用户结构
type User struct {
	Model
	Email             string `gorm:"index:idx_email,unique"`
	EncryptedPassword string
}

//通过邮箱找用户
func FindUserByEmail(email string) (*User, SearchResult) {
	var user User
	result := Result(db.DB.Where("email = ?", email).First(&user).Error)
	return &user, result

}

//通过id找用户
func FindUserByID(id string) (*User, SearchResult) {
	var user User
	result := Result(db.DB.Where("id = ?", id).First(&user).Error)
	return &user, result
}

//通过某一栏找用户
func FindUserByColum(colum string, value interface{}) (*User, bool) {
	var user User
	qs := fmt.Sprintf("%s = ?", colum)
	err := db.DB.Where(qs, value).First(&user).Error
	return &user, errors.Is(err, gorm.ErrRecordNotFound)
}
