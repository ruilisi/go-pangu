package main

import (
	_ "go-jwt/conf"
	"go-jwt/db"
	"go-jwt/redis"
	"go-jwt/routers"
)

func init() {
	db.ConnectDB()
	redis.ConnectRedis()
}

func main() {
	if db.DB != nil {
		defer db.DB.Close()
	}

	routers.InitRouter()
}
