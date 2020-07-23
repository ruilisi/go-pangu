package db

import (
	"fmt"

	"go-jwt/models"
	"go-jwt/setting"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	//replace your database
	DB, err = gorm.Open(setting.DatabaseSetting.Type,
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Port,
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Name,
			setting.DatabaseSetting.Password))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func Close() {
	DB.Close()
}

func FindUserByEmail(email string) models.User {
	var user models.User
	DB.Where("email = ?", email).First(&user)
	return user
}

func FindUserById(id string) models.User {
	var user models.User
	DB.Where("id = ?", id).First(&user)
	return user
}
