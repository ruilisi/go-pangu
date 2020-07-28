package routers

import (
	"go-jwt/conf"
	"go-jwt/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bear := c.Request.Header.Get("Authorization")
		token := strings.Replace(bear, "Bearer ", "", 1)
		sub, err := jwt.Decoder(token)
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, err.Error())
		} else {
			c.Set("sub", sub)
		}
		c.Next()
	}
}

func CheckDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := c.Request.FormValue("DEVICE_TYPE")
		if _, ok := conf.DEVICES[device]; !ok {
			c.Abort()
			c.String(http.StatusBadRequest, "error device")
			return
		}
		c.Next()
	}
}
