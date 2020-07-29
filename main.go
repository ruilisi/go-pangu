package main

import (
	"fmt"
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
	fmt.Println("server starting...")
	routers.InitRouter()
}
