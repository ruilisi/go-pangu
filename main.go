package main

import (
	"go-jwt/conf"
	"go-jwt/db"
	"go-jwt/redis"
	"go-jwt/routers"
)

func init() {
	conf.ReadConf()
	db.ConnectDB()
	redis.ConnectRedis()
}

func main() {
	routers.InitRouter()
}
