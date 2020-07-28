package service

import (
	"go-jwt/args"
	"go-jwt/db"
	"go-jwt/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignInHandler(c *gin.Context) {
	var signIn args.SignIn
	if err := c.ShouldBind(&signIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := signIn.Password
	user := db.FindUserByEmail(signIn.Email)
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout or password error"})
		return
	}
	payload := jwt.GenPayload(signIn.Device, "user", user.ID.String())
	tokenString := jwt.Encoder(payload)
	jwt.OnJwtDispatch(payload)

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}
