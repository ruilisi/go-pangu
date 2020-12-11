package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/i18n/gi18n"
)

var cities = []string{
	"City_SuZhou",
	"City_Peking",
}

func PingHandler(c *gin.Context) {
	//ping接口，测试连通性
	c.String(http.StatusOK, "pong")
}

func CityListHandler(c *gin.Context) {
	t := gi18n.New()
	var list []string
	//绑定数据
	var params map[string]string
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	language := params["language"]
	t.SetLanguage(language)
	for _, city := range cities {
		list = append(list, t.Translate(city))
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"cities_list": list,
	})
}
