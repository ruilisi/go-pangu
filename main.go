package main

import (
	"fmt"
	"go-jwt/db"
	"go-jwt/redis"
	"go-jwt/service"
	"go-jwt/setting"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	setting.Setup()
	db.ConnectDB()
	redis.ConnectRedis()
}

func main() {
	defer db.Close()

	router := gin.Default()
	router.GET("/ping", service.PingHandler)
	router.GET("/auth_ping", service.AuthPingHandler)
	router.POST("/sign_in", service.SignInHandler)
	router.POST("/change_password", service.ChangePasswordHandler)
	router.Run(fmt.Sprintf(":%v", setting.ServerSetting.HttpPort))
}
