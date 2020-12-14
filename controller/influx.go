package controller

import (
	"fmt"
	"go-pangu/influx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveInfluxDBHandler(c *gin.Context) {
	//绑定数据
	var params map[string]string
	var points []influx.Point
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := influx.UserInfo{InfluxMeasurement: "userinfo", UserName: params["user_name"], Local: params["local"], Version: params["version"]}

	point := influx.Point{
		Struct: user,
	}

	points = append(points, point)
	influx.WritePoints(points)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func ShowInfuxDBHandler(c *gin.Context) {
	fmt.Println(12312312)
	points := influx.ReadPoints(`select * from userinfo limit 1000`, []influx.UserInfoRead{}).([]influx.UserInfoRead)
	fmt.Println(points)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"points": points,
	})
}
