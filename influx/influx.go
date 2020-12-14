package influx

import (
	"fmt"
	"go-pangu/conf"
	"log"

	"github.com/cbrake/influxdbhelper/v2"
	client "github.com/influxdata/influxdb1-client/v2"
)

var influxURL = fmt.Sprintf("http://%s:%s", conf.GetEnv("INFLUXDB_HOST"), conf.GetEnv("INFLUXDB_PORT"))
var db = conf.GetEnv("INFLUXDB_DB")
var c influxdbhelper.Client

func ConnectInflux() {
	c, _ = influxdbhelper.NewClient(influxURL, "", "", "ns")
}

// Init initializes the database connection
func Init() (err error) {
	// Create test database if it doesn't already exist
	q := client.NewQuery("CREATE DATABASE "+db, "", "")
	res, err := c.Query(q)
	if err != nil {
		return err
	}
	if res.Error() != nil {
		log.Println("dbhelper db initialize failed")
		return res.Error()
	}

	return nil
}

// write data to database
func WritePoints(points []interface{}) {
	c = c.UseDB(db)
	for _, p := range points {
		err := c.WritePoint(p)
		if err != nil {
			log.Fatal("Error writing point: ", err)
		}
	}
}

// query data from db
func ReadPoints(query string, points interface{}) interface{} {
	err := c.UseDB(db).DecodeQuery(query, &points)
	if err != nil {
		log.Fatal("Query error: ", err)
	}

	return points
}
