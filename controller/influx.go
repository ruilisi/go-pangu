package controller

import (
	"go-pangu/influx"
	"net/http"

	"github.com/gin-gonic/gin"
)

// if you want use this file apis, please uncomment main.go two lines code and install influxdb
// influx.ConnectInflux()  and  influx.Init()
// before use, please run  main.go -db=create  to create influxdb database
func SaveInfluxDBHandler(c *gin.Context) {
	//绑定数据
	var params map[string]string
	var points []interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := influx.UserInfo{InfluxMeasurement: "userinfo", UserName: params["user_name"],
		Local: params["local"], Version: params["version"]}

	points = append(points, user)
	influx.WritePoints(points)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func ShowInfuxDBHandler(c *gin.Context) {
	var points []interface{}
	influx.ReadPoints(`select * from userinfo limit 1000`, &points)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"points": points,
	})
}
