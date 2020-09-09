package service

import (
	"go-pangu/conf"
	"go-pangu/db"
	"go-pangu/jwt"
	"go-pangu/models"
	"go-pangu/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context) {
	var params map[string]string
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password := params["password"]
	password_confirm := params["password_confirm"]
	signupType := params["signup_type"]
	var user *models.User
	var notFound bool
	switch signupType {
	case "email":
		user, notFound = models.FindUserByColum("email", params["email"])
	default:
		notFound = false
	}
	if !notFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "aacount exists"})
		return
	}

	if password != password_confirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "password confirmation dismatch"})
		return
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	user = &models.User{Email: params["email"], EncryptedPassword: string(bcryptedPassword)}
	db.DB.Create(user)
	c.JSON(http.StatusOK, gin.H{"status": "register success"})
}

func SignInHandler(c *gin.Context) {
	var params map[string]string
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := params["password"]
	deviceType := params["DEVICE_TYPE"]
	loginType := params["login_type"]

	if !util.Contains(conf.DEVICE_TYPES, deviceType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "device type error"})
		return
	}
	var user *models.User
	var notFound bool
	switch loginType {
	case "email":
		user, notFound = models.FindUserByColum("email", params["email"])
	default:
		notFound = true
	}
	if notFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "accout or password error"})
		return
	}
	payload := jwt.GenPayload(deviceType, "user", user.ID.String())
	tokenString := jwt.Encoder(payload)
	jwt.OnJwtDispatch(payload)

	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}
