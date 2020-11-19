package controller

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
	//需要验证token的ping，需要先登录才能使用
	c.String(http.StatusOK, fmt.Sprintf("pong"))
}

func ChangePasswordHandler(c *gin.Context) {
	//修改密码的handler
	sub, _ := c.Get("sub")
	scp, _ := c.Get("scp")
	var change params.ChangePassword
	var oldEncryptedPassword string
	var user *models.User

	//绑定参数 这里用的是结构体绑定
	if err := c.ShouldBind(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if change.Password != change.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "password and password confirm not match"})
		return
	}

	//通过scp判断不同类型用户 这里可以自行添加不同类型用户情况
	//这里查找旧密码，与传入的旧密码加密后的字符串比较是否一致
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

	//将新密码加密后保存
	hash, err := bcrypt.GenerateFromPassword([]byte(change.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//更新数据库 生成新token  注销老token
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

	//返回结果
	c.JSON(http.StatusOK, gin.H{"status": "update password success"})
}
