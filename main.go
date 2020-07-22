package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type Payload struct {
	Device string `json:"device,omitempty"`
	Scp    string `json:"scp,omitempty"`
	jwt.StandardClaims
}

type User struct {
	Id                string
	Email             string `form:"email" json:"email" xml:"email" binding:"required"`
	Password          string `form:"password" json:"password" xml:"password" binding:"required"`
	EncryptedPassword string
	Device            string `form:"DEVICE_TYPE" json:"DEVICE_TYPE" xml:"DEVICE_TYPE" binding:"required"`
	Type              string `form:"type" json:"type" xml:"type" binding:"required"`
}

var hmacSampleSecret = "RANDOM_SECRET"
var ctx = context.Background()
var db *gorm.DB
var rdb *redis.Client

func ConnectDB() {
	var err error
	//replace your database
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=chengzi_development password=postgres")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func init() {
	ConnectDB()
	ConnectRedis()
}

func main() {
	defer db.Close()

	router := gin.Default()
	router.GET("/ping", PingHandler)
	router.GET("/auth_ping", AuthPingHandler)
	router.POST("/sign_in", SignInHandler)
	router.Run(":3000")
}

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func AuthPingHandler(c *gin.Context) {
	bear := c.Request.Header.Get("Authorization")
	token := strings.Replace(bear, "Bear ", "", 1)
	_, err := Decoder(token)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	c.String(http.StatusOK, "auth pong")
}

func FindUserByEmail(email string) User {
	var user User
	db.Where("email = ?", email).First(&user)
	return user
}

func GenPayload(device, scp, sub string) Payload {
	now := time.Now()
	return Payload{
		Device: device,
		Scp:    scp,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(1 * time.Hour).Unix(),
			Id:        uuid.New().String(),
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   sub,
		},
	}
}

func SignInHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device := user.Device
	password := user.Password
	user = FindUserByEmail(user.Email)
	if user.Email == "" {
		c.JSON(http.StatusOK, gin.H{"status": "accout not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "accout or password error"})
		return
	}
	payload := GenPayload(device, "user", user.Id)
	tokenString := Encoder(payload)
	OnJwtDispatch(user, payload)

	c.Header("Authorization", "Bear "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}

func JwtRevoked(payload Payload) bool {
	ok, _ := rdb.Exists(ctx, fmt.Sprintf("user_blacklist:%s:%s:%s", payload.Subject, payload.Device, payload.Id)).Result()
	return ok == 1
}

func RevokeJwt(payload Payload, user User) {
	expiration := payload.ExpiresAt - payload.IssuedAt
	rdb.Set(ctx, fmt.Sprintf("user_blacklist:%s:%s:%s", user.Id, payload.Device, payload.Id), payload.Id, time.Duration(expiration)*time.Second)
}

func OnJwtDispatch(user User, payload Payload) {
	iat := time.Now()
	lastJwt, err := rdb.Get(ctx, fmt.Sprintf("user_device_jwt:%s:%s", user.Id, payload.Device)).Result()
	if err == redis.Nil {
		fmt.Println("redis key not found")
	}
	if lastJwt != "" {
		arr := strings.Split(lastJwt, ":")
		jti, expStr := arr[0], arr[len(arr)-1]
		exp, err := strconv.ParseInt(expStr, 10, 64)
		if err != nil {
			exp = time.Now().Unix()
		}
		newPayload := payload
		newPayload.Id = jti
		newPayload.IssuedAt = time.Now().Unix()
		newPayload.ExpiresAt = exp
		RevokeJwt(newPayload, user)
	}

	rdb.Set(ctx, fmt.Sprintf("user_device_jwt:%s:%s", user.Id, payload.Device), fmt.Sprintf("%s:%d", payload.Id, payload.ExpiresAt), time.Unix(payload.ExpiresAt, 0).Sub(iat))
}

func Encoder(payload Payload) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func Decoder(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		return "", err
	}

	if payload, ok := token.Claims.(*Payload); ok && token.Valid {
		sub := (*payload).Subject
		if sub != "" && !JwtRevoked(*payload) {
			return sub, nil
		} else {
			return "", fmt.Errorf("token is expired")
		}
	}

	return "", fmt.Errorf("invalid token")
}
