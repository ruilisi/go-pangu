package service

import (
	"go-pangu/db"
	"go-pangu/jwt"
	"go-pangu/models"
	"go-pangu/params"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context) {
	var signUp params.SignUp
	if err := c.ShouldBind(&signUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.FindUserByEmail(signUp.Email)
	if user.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "accout exists"})
		return
	}

	if signUp.Password != signUp.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "password confirmation dismatch"})
		return
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(signUp.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	user = &models.User{Email: signUp.Email, EncryptedPassword: string(bcryptedPassword)}
	db.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{"status": "register success"})
}

func SignInHandler(c *gin.Context) {
	var signIn params.SignIn
	if err := c.ShouldBind(&signIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := signIn.Password
	user := models.FindUserByEmail(signIn.Email)
	if user.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout or password error"})
		return
	}
	payload := jwt.GenPayload("user", user.ID.String())
	tokenString := jwt.Encoder(payload)
	jwt.OnJwtDispatch(payload)

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}
