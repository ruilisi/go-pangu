package main

import (
	"fmt"
	"go-pangu/args"
	"go-pangu/db"
	"go-pangu/models"
	"go-pangu/redis"
	"go-pangu/routers"
	"log"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/judwhite/go-svc/svc"
)

// func init() {
// redis.ConnectRedis()
// }

// func main() {
// args.ParseCmd()
// switch args.Cmd.DB {
// case "create":
// db.Create()
// case "migrate":
// db.Migrate(args.Cmd.GIN_ENV, &models.User{})
// case "seed":
// default:
// ctx, cancel := context.WithCancel(context.Background())
// fmt.Println("111111")
// osSignal := make(chan os.Signal)
// signal.Notify(osSignal, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)
// fmt.Println("server starting...")
// db.Open("")
// defer db.Close()
// routers.InitRouter(ctx, cancel, osSignal)
// }
// }
type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

func (p *program) Init(env svc.Environment) error {
	redis.ConnectRedis()
	return nil
}

func (p *program) Start() error {
	args.ParseCmd()
	switch args.Cmd.DB {
	case "create":
		fmt.Println("creating database")
		db.Create()
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	case "migrate":
		fmt.Println("migrating tables")
		db.Migrate(args.Cmd.GIN_ENV, &models.User{})
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	case "seed":
	case "drop":
		fmt.Println("droping database")
		if args.Cmd.TABLE != "" {
			db.Open("")
			db.DB.Migrator().DropTable(args.Cmd.TABLE)
		} else {
			db.Drop()
		}
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	default:
		fmt.Println("server starting...")
		db.Open("")
		routers.InitRouter(os.Interrupt)
	}
	return nil
}

func (p *program) Stop() error {
	fmt.Println("\nserver stoping")
	time.Sleep(time.Duration(1) * time.Second)
	return nil
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, os.Interrupt); err != nil {
		log.Fatal(err)
	}
}
