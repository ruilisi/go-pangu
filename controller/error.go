package controller

import "github.com/gin-gonic/gin"

func StatusError(c *gin.Context, httpcode int, status string, err string) {
	//错误返回的一个模板，可以省去一些重复劳动
	//httpcode是http.Status...形式，也就是int值(例如StatusOK是400)
	c.JSON(httpcode, gin.H{
		"status": status,
		"error":  err,
	})
}
