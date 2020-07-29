package db

import (
	"go-jwt/conf"
	"go-jwt/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type PGExtension struct {
	Extname string
}

func openDB() (err error) {
	url := conf.GetEnv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	var pgExtension PGExtension
	DB.Table("pg_extension").Where("extname = ?", "pgcrypto").Find(&pgExtension)
	if pgExtension.Extname != "pgcrypto" {
		DB.Exec("CREATE EXTENSION pgcrypto")
	}
	return
}

func ConnectDB() {
	if err := openDB(); err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&models.User{})
	email := "test@123.com"
	user := FindUserByEmail(email)
	if user.Email == "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
		user := models.User{Email: email, EncryptedPassword: string(hash)}
		DB.Create(&user)
	}
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
