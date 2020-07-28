package db

import (
	"go-jwt/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func CreateDB() error {
// cmd := exec.Command("createdb", viper.Get("DATABASE_URL").(string))
// return cmd.Run()
// }

func openDB() (err error) {
	url := viper.Get("DATABASE_URL").(string)
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	return
}

func ConnectDB() {
	if err := openDB(); err != nil {
		panic(err.Error())
	}

	// DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
	// user := models.User{Email: "Jinzhu12"}
	// DB.Create(&user)
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
