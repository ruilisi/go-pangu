package args

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Param(c *gin.Context, key string) string {
	if c.ContentType() == binding.MIMEJSON {
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		var sec map[string]string
		if err := json.Unmarshal(bodyBytes, &sec); err == nil {
			v, exist := sec[key]
			if exist {
				return v
			}
		}
	}
	return c.Request.FormValue(key)
}

func Params(c *gin.Context) map[string]string {
	sec := make(map[string]string)
	if c.ContentType() == binding.MIMEJSON {
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		if err := json.Unmarshal(bodyBytes, &sec); err == nil {
			return sec
		}
	}
	for k, v := range c.Request.URL.Query() {
		sec[k] = v[0]
	}
	return sec
}

type SignUp struct {
	Email           string `form:"email" json:"email" xml:"email" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" xml:"password_confirm" binding:"required"`
}

type SignIn struct {
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type ChangePassword struct {
	OriginPassword  string `form:"origin_password" json:"origin_password" xml:"origin_password" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" xml:"password_confirm" binding:"required"`
}
