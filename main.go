package main

import (
	"go-jwt/db"
	"go-jwt/redis"
	"go-jwt/routers"
	"go-jwt/setting"
)

func init() {
	setting.Setup()
	db.ConnectDB()
	redis.ConnectRedis()
}

func main() {
	defer db.Close()

	routers.InitRouter()
}
