package influx

import (
	"time"

	"github.com/cbrake/influxdbhelper/v2"
)

type Point struct {
	Struct interface{}
}

type PointRead struct {
	Struct interface{}
}

type UserInfo struct {
	InfluxMeasurement influxdbhelper.Measurement
	UserName          string `influx:"user_name,tag"`
	Local             string `influx:"local"`
	Version           string `influx:"version"`
}

type UserInfoRead struct {
	InfluxMeasurement influxdbhelper.Measurement
	Time              time.Time `influx:"time"`
	UserName          string    `influx:"user_name,tag"`
	Local             string    `influx:"local"`
	Version           string    `influx:"version"`
}
