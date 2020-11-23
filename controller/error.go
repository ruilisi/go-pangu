package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatusError(c *gin.Context, httpcode int, status string, err string) {
	//错误返回的一个模板，可以省去一些重复劳动
	//httpcode是http.Status...形式，也就是int值(例如StatusOK是400)
	c.JSON(httpcode, gin.H{
		"status": status,
		"error":  err,
	})
}

func Select(c *gin.Context, tx *gorm.DB, ch chan string, finish chan int, resp map[string]interface{}) {
	select {
	case status := <-ch:
		errmsg := <-ch
		tx.Rollback()
		StatusError(c, http.StatusBadRequest, status, errmsg)
	case <-finish:
		tx.Commit()
		c.JSON(http.StatusOK, gin.H(resp))
	case <-time.After(5 * time.Second):
		tx.Rollback()
		StatusError(c, http.StatusBadRequest, "fail", "time out")
	}
}
