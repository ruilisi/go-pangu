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

func Open(env string) {
	var err error
	var url string
	url = conf.GetEnv("DATABASE_URL")
	if env == "test" {
		url = conf.GetEnv("DATABASE_TESTURL")
	}
	if DB, err = gorm.Open(postgres.Open(url), &gorm.Config{}); err != nil {
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
		if conf.GetEnv("GIN_ENV") != "production" {
			baseDb.Exec(fmt.Sprintf("CREATE DATABASE %s;", "go_pangu_test"))
		}
	}
}

func Migrate(env string, models ...interface{}) {
	Open(env)
	var pgExtension PGExtension
	DB.Table("pg_extension").Where("extname = ?", "pgcrypto").Find(&pgExtension)
	if pgExtension.Extname != "pgcrypto" {
		DB.Exec("CREATE EXTENSION pgcrypto")
	}
	if env == "test" {
		DB.Exec(`CREATE OR REPLACE FUNCTION truncate_tables(username IN VARCHAR) RETURNS void AS $$
DECLARE
		statements CURSOR FOR 
				SELECT tablename FROM pg_tables
				WHERE tableowner = username AND schemaname = 'public';
BEGIN 
		FOR stmt IN statements LOOP 
				EXECUTE 'TRUNCATE TABLE ' || quote_ident(stmt.tablename) || ' CASCADE;';  
		END LOOP; 
END; 
$$ LANGUAGE plpgsql;`)
	}
	DB.AutoMigrate(models...)
}

func CleanTablesData() {
	Open("test")
	DB.Exec(`SELECT truncate_tables('postgres');`)
}

func DropTables(env string) {
	Open(env)
	DB.Exec("DROP SCHEMA public CASCADE;")
	defer Close()
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
		if conf.GetEnv("GIN_ENV") != "production" {
			baseDb.Exec(fmt.Sprintf("DROP DATABASE %s;", "go_pangu_test"))
		}
	}
}

func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
