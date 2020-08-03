package db

import (
	"fmt"
	"go-jwt/conf"
	"go-jwt/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type PGExtension struct {
	Extname string
}

func openDB() {
	var err error
	url := conf.GetEnv("DATABASE_URL")
	base_url := conf.GetEnv("BASE_DATABASE_URL")

	if DB, err = gorm.Open(postgres.Open(url), &gorm.Config{}); err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			baseDb, err := gorm.Open(postgres.Open(base_url), &gorm.Config{})
			if err != nil {
				panic(err.Error())
			}
			baseDb = baseDb.Exec(fmt.Sprintf("CREATE DATABASE %s;", conf.GetEnv("DATABASE_NAME")))
			sqlDB, err := baseDb.DB()
			sqlDB.Close()

			DB, err = gorm.Open(postgres.Open(conf.GetEnv("DATABASE_URL")), &gorm.Config{})
			if err != nil {
				panic(err.Error())
			}
		} else {
			panic(err.Error())
		}
	}

	var pgExtension PGExtension
	DB.Table("pg_extension").Where("extname = ?", "pgcrypto").Find(&pgExtension)
	if pgExtension.Extname != "pgcrypto" {
		DB.Exec("CREATE EXTENSION pgcrypto")
	}
}

func ConnectDB() {
	openDB()

	DB.AutoMigrate(&models.User{})
	email := "test@123.com"
	user := FindUserByEmail(email)
	if user.Email == "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
		user = models.User{Email: email, EncryptedPassword: string(hash)}
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
