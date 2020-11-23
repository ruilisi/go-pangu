package controller

import (
	"go-pangu/conf"
	"go-pangu/db"
	"go-pangu/jwt"
	"go-pangu/models"
	"go-pangu/util"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(c *gin.Context) {
	//注册user的handler

	//绑定数据
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "account exists"})
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

//创建十个账户，如果其中一个创建失败，一起失败。整合了并发跟数据库回滚和超时。
//create ten users,if one fail,both fail.(with goroutine,database rollback,time out)
func CreateUsersHandler(c *gin.Context) {
	var (
		params map[string]interface{}
	)

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := params["password"].(string)

	ch := make(chan string, 2) //get error message
	finish := make(chan int)   // get finish signal
	var wg sync.WaitGroup
	tx := db.DB.Begin()
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				ch <- "fail"
				ch <- err.Error()
				return
			}
			user1 := &models.User{Email: RandStringRunes(6, LetterRunes) + params["email"].(string),
				EncryptedPassword: string(bcryptedPassword)}
			err = tx.Create(user1).Error
			if err != nil {
				ch <- "fail"
				ch <- err.Error()
				return
			}
			defer wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		finish <- 1
	}()

	resp := map[string]interface{}{
		"status": "success",
	}
	Select(c, tx, ch, finish, resp)
}
