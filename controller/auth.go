package service

import (
	"fmt"
	"go-pangu/conf"
	"go-pangu/db"
	"go-pangu/jwt"
	"go-pangu/models"

	//"go-pangu/models/user"
	"go-pangu/params"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CurrentUser(c *gin.Context) *models.User {
	sub, _ := c.Get("sub")
	user, _ := models.FindUserByColum("id", sub)
	return user
}

func AuthPingHandler(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("pong"))
}

func ChangePasswordHandler(c *gin.Context) {
	sub, _ := c.Get("sub")
	scp, _ := c.Get("scp")
	var change params.ChangePassword
	var oldEncryptedPassword string
	var user *models.User
	if err := c.ShouldBind(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if change.Password != change.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "password and password confirm not match"})
		return
	}

	switch scp {
	case "user":
		user, _ = models.FindUserByColum("id", sub)
		oldEncryptedPassword = user.EncryptedPassword

	}
	err := bcrypt.CompareHashAndPassword([]byte(oldEncryptedPassword), []byte(change.OriginPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "origin password error"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(change.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encryptedPassword := string(hash)
	var payload jwt.Payload
	switch scp {
	case "user":
		db.DB.Model(&user).Updates(models.User{EncryptedPassword: encryptedPassword})
		payload = jwt.GenPayload("", "user", user.ID.String())
		for _, device := range conf.DEVICE_TYPES {
			payload.Device = device
			jwt.RevokeLastJwt(payload)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "update password success"})
}
