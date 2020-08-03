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
	authorized := router.Group("/")
	authorized.Use(Auth())
	{
		authorized.GET("/auth_ping", service.AuthPingHandler)
	}

	users := router.Group("/users")
	{
		users.POST("/sign_up", service.SignUpHandler)
		users.POST("/sign_in", service.SignInHandler)
	}
	users.Use(Auth())
	{
		users.POST("/change_password", service.ChangePasswordHandler)
	}
	router.Run(fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")))
}
