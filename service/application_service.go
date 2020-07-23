package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var DEVICES = map[string]bool{"WINDOWS": true, "MAC": true, "ANDROID": true, "IOS": true}

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
