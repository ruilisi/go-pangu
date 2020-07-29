package routers

import (
	"bytes"
	"encoding/json"
	"go-jwt/conf"
	"go-jwt/jwt"
	"io/ioutil"
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
		device := c.Request.FormValue("device_type")
		if _, ok := conf.DEVICES[device]; ok {
			c.Next()
			return
		}

		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		var sec map[string]string
		if err := json.Unmarshal(bodyBytes, &sec); err == nil {
			if _, ok := conf.DEVICES[sec["DEVICE_TYPE"]]; ok {
				c.Next()
				return
			}
		}

		c.Abort()
		c.String(http.StatusBadRequest, "error device")
	}
}
