package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	//ping接口，测试连通性
	c.String(http.StatusOK, "pong")
}
