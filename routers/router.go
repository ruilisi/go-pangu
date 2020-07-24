package routers

import (
	"fmt"

	"go-jwt/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/ping", service.PingHandler)
	router.POST("/sign_in", service.SignInHandler)
	router.Use(Auth())
	{
		router.GET("/auth_ping", service.AuthPingHandler)
		router.POST("/change_password", service.ChangePasswordHandler)
	}
	router.Run(fmt.Sprintf(":%v", viper.Get("HTTP_PORT")))
}
