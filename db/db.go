package db

import (
	"fmt"
	"go-jwt/conf"
	"go-jwt/models"
	"net/url"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type PGExtension struct {
	Extname string
}

func Open() {
	var err error
	if DB, err = gorm.Open(postgres.Open(conf.GetEnv("DATABASE_URL")), &gorm.Config{}); err != nil {
		panic(err.Error())
	}
}

func Create() {
	dbURL := conf.GetEnv("DATABASE_URL")
	if uri, err := url.Parse(dbURL); err != nil {
		panic(err)
	} else {
		path := uri.Path
		uri.Path = ""
		baseDb, err := gorm.Open(postgres.Open(uri.String()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		baseDb.Exec(fmt.Sprintf("CREATE DATABASE %s;", path[1:]))
		uri.Path = path
	}

	Open()
	var pgExtension PGExtension
	DB.Table("pg_extension").Where("extname = ?", "pgcrypto").Find(&pgExtension)
	if pgExtension.Extname != "pgcrypto" {
		DB.Exec("CREATE EXTENSION pgcrypto")
	}
}

func Seed() {
	Open()
	email := "test@123.com"
	user := FindUserByEmail(email)
	if user.Email == "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
		user = &models.User{Email: email, EncryptedPassword: string(hash)}
		DB.Create(user)
	}
}

func Migrate() {
	Open()
	DB.AutoMigrate(&models.User{})
}

func Drop() {
	dbURL := conf.GetEnv("DATABASE_URL")
	if uri, err := url.Parse(dbURL); err != nil {
		panic(err)
	} else {
		path := uri.Path
		uri.Path = ""
		baseDb, err := gorm.Open(postgres.Open(uri.String()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		baseDb.Exec(fmt.Sprintf("DROP DATABASE %s;", path[1:]))
	}
}

func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}

func FindUserByEmail(email string) *models.User {
	var user models.User
	DB.Where("email = ?", email).First(&user)
	return &user
}

func FindUserById(id string) *models.User {
	var user models.User
	DB.Where("id = ?", id).First(&user)
	return &user
}
