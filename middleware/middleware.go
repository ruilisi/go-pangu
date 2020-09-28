package middleware

import (
	"go-pangu/controller"
	"go-pangu/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(scp string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bear := c.Request.Header.Get("Authorization")
		token := strings.Replace(bear, "Bearer ", "", 1)
		sub, scope, err := jwt.Decoder(token)
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, err.Error())
		} else {
			if scope != scp {
				controller.StatusError(c, http.StatusUnauthorized, "unauthorized", "invalid scope")
				c.Abort()
			}
			c.Set("sub", sub)
			c.Set("scp", scope)
			c.Next()
		}
	}
}
