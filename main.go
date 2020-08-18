package main

import (
	"fmt"
	"go-pangu/args"
	"go-pangu/db"
	"go-pangu/models"
	"go-pangu/redis"
	"go-pangu/routers"
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
		models.Seed()
	default:
		fmt.Println("server starting...")
		db.Open()
		defer db.Close()
		routers.InitRouter()
	}
}
