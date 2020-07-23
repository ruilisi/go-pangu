package service

import (
	"go-jwt/db"
	"go-jwt/jwt"
	"go-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device := user.Device
	if _, ok := DEVICES[device]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error device"})
		return
	}

	password := user.Password
	user = db.FindUserByEmail(user.Email)
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout or password error"})
		return
	}
	payload := jwt.GenPayload(device, "user", user.Id)
	tokenString := jwt.Encoder(payload)
	jwt.OnJwtDispatch(payload)

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}
