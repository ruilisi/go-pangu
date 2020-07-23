package main

import (
	"go-jwt/db"
	"go-jwt/service"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	db.ConnectDB()
	db.ConnectRedis()
}

func main() {
	defer db.DB.Close()

	router := gin.Default()
	router.GET("/ping", service.PingHandler)
	router.GET("/auth_ping", service.AuthPingHandler)
	router.POST("/sign_in", service.SignInHandler)
	router.POST("/change_password", service.ChangePasswordHandler)
	router.Run(":3000")
}
