package main

import (
	"fmt"
	"go-jwt/args"
	_ "go-jwt/conf"
	"go-jwt/db"
	"go-jwt/redis"
	"go-jwt/routers"
)

func init() {
	redis.ConnectRedis()
}

func main() {
	args.ParseCmd()
	switch args.Cmd.DB {
	case "create":
		db.Create()
	case "migrate":
		db.Migrate()
	case "seed":
		db.Seed()
	default:
		fmt.Println("server starting...")
		routers.InitRouter()
	}
}
