package db

import (
	"fmt"
	"go-pangu/conf"
	"net/url"

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

func Migrate(models ...interface{}) {
	Open()
	DB.AutoMigrate(models...)
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
