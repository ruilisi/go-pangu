package controller

import (
	"go-pangu/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/i18n/gi18n"
)

var cities = []string{
	"City_SuZhou",
	"City_Peking",
}

// WSAuthedUser broadcasts event to authed user
func WSAuthedUser(typ string, data map[string]interface{}) {
	hub := websocket.GetHub()
	data["type"] = typ
	hub.SysBroadcastJSON("websocketTest", data)
}

func wsTest() {
	WSAuthedUser("test", map[string]interface{}{})
}

func PingHandler(c *gin.Context) {
	//ping接口，测试连通性
	wsTest()
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
