package routers

import (
	"fmt"

	"go-jwt/conf"
	"go-jwt/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.GET("/ping", service.PingHandler)
	router.POST("/sign_up", service.SignUpHandler)
	router.POST("/sign_in", CheckDevice(), service.SignInHandler)
	authorized := router.Group("/")
	authorized.Use(Auth())
	{
		authorized.GET("/auth_ping", service.AuthPingHandler)
		authorized.POST("/change_password", service.ChangePasswordHandler)
	}
	router.Run(fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")))
}
