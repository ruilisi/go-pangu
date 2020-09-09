package routers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-pangu/conf"
	service "go-pangu/controller"
	"go-pangu/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(ctx context.Context, cancel context.CancelFunc, osSignal chan os.Signal) {
	router := SetupRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.GetEnv("HTTP_PORT")),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	fmt.Printf("Listening and serving HTTP on :%s\n", conf.GetEnv("HTTP_PORT"))
	<-osSignal
	cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Serve forced to shutdown:", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Server exiting")
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
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/auth_ping", service.AuthPingHandler)
	}
	users := router.Group("/users")
	{
		users.POST("/sign_up", service.SignUpHandler)
		users.POST("/sign_in", service.SignInHandler)
	}
	users.Use(middleware.Auth())
	{
		//users.POST("/change_password", service.ChangePasswordHandler)
	}
	return router

}
