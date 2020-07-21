package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Payload struct {
	Device string `json:"device,omitempty"`
	Scp    string `json:"scp,omitempty"`
	jwt.StandardClaims
}

//TODO encrypt password
type User struct {
	Id       string
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
	Device   string `form:"DEVICE_TYPE" json:"DEVICE_TYPE" xml:"DEVICE_TYPE"  binding:"required"`
	Type     string `form:"type" json:"type" xml:"type"  binding:"required"`
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

func SignInHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO: check password in db
	//generate Token
	tokenString := Encoder(user.Device)
	//TODO OnJwtDispatch()
	c.Header("Authorization", "Bear "+tokenString)
	c.JSON(http.StatusOK, gin.H{"status": "login success"})
}

func JwtRevoked(payload Payload, user User) bool {
	ok, _ := rdb.Exists(ctx, fmt.Sprintf("user_blacklist:%s:%s:%s", user.Id, payload.Device, payload.Id)).Result()
	return ok == 1
}

func RevokeJwt(payload Payload, user User) {
	expiration := payload.ExpiresAt - payload.IssuedAt
	rdb.Set(ctx, fmt.Sprintf("user_blacklist:%s:%s:%s", user.Id, payload.Device, payload.Id), payload.Id, time.Duration(expiration)*time.Second)
}

func OnJwtDispatch(user User, payload Payload) {
	iat := time.Now()
	lastJwt, err := rdb.Get(ctx, fmt.Sprintf("Jwt:%s:%s", user.Id, payload.Device)).Result()
	if err == redis.Nil {
		fmt.Println("redis key not found")
	}
	if lastJwt != "" {
		arr := strings.Split(lastJwt, ":")
		jti, exp := arr[0], arr[len(arr)-1]
		fmt.Println(jti, exp)
		RevokeJwt(payload, user)
	}

	rdb.Set(ctx, fmt.Sprintf("user_device_jwt:%s:%s", user.Id, payload.Device), fmt.Sprintf("%s:%d", payload.Id, payload.ExpiresAt), time.Unix(payload.ExpiresAt, 0).Sub(iat))
}

func Encoder(device string) string {
	now := time.Now()
	claims := Payload{
		Device: device,
		Scp:    "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(1 * time.Hour).Unix(),
			Id:        uuid.New().String(),
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   "f955e4fe-723c-4f1e-88bb-12df4d1bdb55",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	fmt.Println(tokenString, err)
	return tokenString
}

func Decoder(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, _ := claims["sub"]
		return sub.(string), nil
	} else {
		return "", err
	}
}
