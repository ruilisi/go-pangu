package main

import (
	"context"
	"fmt"
	"go-pangu/args"
	"go-pangu/db"
	"go-pangu/models"
	"go-pangu/redis"
	"go-pangu/routers"
	"os"
	"os/signal"
	"syscall"
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
		db.Migrate(args.Cmd.GIN_ENV, &models.User{})
	case "seed":
	//	models.Seed()
	default:
		ctx, cancel := context.WithCancel(context.Background())
		fmt.Println("111111")
		osSignal := make(chan os.Signal)
		signal.Notify(osSignal, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)
		fmt.Println("server starting...")
		db.Open("")
		defer db.Close()
		//	if conf.GetEnv("GIN_ENV") == "production" {
		//	rabbitmq.ConsumeM("go_pangu.CollectIpWorker", ctx, args.Cmd.Amqp)
		//}
		routers.InitRouter(ctx, cancel, osSignal)
	}
}
