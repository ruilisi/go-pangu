package routers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go-pangu/conf"
	service "go-pangu/controller"
	"go-pangu/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(sig ...os.Signal) {
	router := SetupRouter()

	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	signalChan := make(chan os.Signal, 1)

	go func() {
		router.Run(fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")))
	}()
	signal.Notify(signalChan, sig...)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.ExposeHeaders = []string{"Authorization"}
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	router.GET("/ping", service.PingHandler)
	//	router.GET("")
	authorized := router.Group("/")
	authorized.Use(middleware.Auth("user"))
	{
		authorized.GET("/auth_ping", service.AuthPingHandler)
	}
	users := router.Group("/users")
	{
		users.POST("/sign_up", service.SignUpHandler)
		users.POST("/sign_in", service.SignInHandler)
	}
	users.Use(middleware.Auth("user"))
	{
		users.POST("/change_password", service.ChangePasswordHandler)
	}
	return router
}
