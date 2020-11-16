package controller

import "github.com/gin-gonic/gin"

func StatusError(c *gin.Context, httpcode int, status string, err string) {
	c.JSON(httpcode, gin.H{
		"status": status,
		"error":  err,
	})
}
