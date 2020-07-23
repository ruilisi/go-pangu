package main

import (
	"fmt"
	"net/http"
	"strings"

	"go-jwt/db"
	"go-jwt/jwt"
	"go-jwt/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type ChangePassword struct {
	OriginPassword  string `form:"origin_password" json:"origin_password" xml:"origin_password" binding:"required"`
	Password        string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" xml:"password_confirm" binding:"required"`
}

var DEVICES = map[string]bool{"WINDOWS": true, "MAC": true, "ANDROID": true, "IOS": true}

func init() {
	db.ConnectDB()
	db.ConnectRedis()
}

func main() {
	defer db.DB.Close()

	router := gin.Default()
	router.GET("/ping", PingHandler)
	router.GET("/auth_ping", AuthPingHandler)
	router.POST("/sign_in", SignInHandler)
	router.POST("/change_password", ChangePasswordHandler)
	router.Run(":3000")
}

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func AuthPingHandler(c *gin.Context) {
	_, err := Auth(c)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	c.String(http.StatusOK, "auth pong")
}

func Auth(c *gin.Context) (string, error) {
	bear := c.Request.Header.Get("Authorization")
	token := strings.Replace(bear, "Bearer ", "", 1)
	sub, err := jwt.Decoder(token)
	if err != nil {
		return "", err
	}
	return sub, nil
}

func SignInHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device := user.Device
	if _, ok := DEVICES[device]; !ok {
		c.JSON(http.StatusOK, gin.H{"status": "error device"})
		return
	}

	password := user.Password
	user = db.FindUserByEmail(user.Email)
	if user.Email == "" {
		c.JSON(http.StatusOK, gin.H{"status": "accout not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "accout or password error"})
		return
	}
	payload := jwt.GenPayload(device, "user", user.Id)
	tokenString := jwt.Encoder(payload)
	jwt.OnJwtDispatch(payload)

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}

func ChangePasswordHandler(c *gin.Context) {
	sub, err := Auth(c)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	var change ChangePassword
	if err := c.ShouldBind(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if change.Password != change.PasswordConfirm {
		c.JSON(http.StatusOK, gin.H{"status": "password and password confirm not match"})
		return
	}

	user := db.FindUserById(sub)
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(change.OriginPassword))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "origin password error"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(change.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encryptedPassword := string(hash)
	db.DB.Model(&user).Updates(models.User{EncryptedPassword: encryptedPassword})
	payload := jwt.GenPayload("", "user", user.Id)
	for device, _ := range DEVICES {
		payload.Device = device
		jwt.RevokeLastJwt(payload)
	}
	c.JSON(http.StatusOK, gin.H{"status": "update password success"})
}
