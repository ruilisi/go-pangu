package db

import (
	"context"
	"fmt"
	"time"

	"go-jwt/models"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
var ctx = context.Background()

func ConnectDB() {
	var err error
	//replace your database
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=lingti_development password=postgres")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Close() {
	DB.Close()
}

func FindUserByEmail(email string) models.User {
	var user models.User
	DB.Where("email = ?", email).First(&user)
	return user
}

func FindUserById(id string) models.User {
	var user models.User
	DB.Where("id = ?", id).First(&user)
	return user
}

func RDBGet(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func RDBSet(key string, value interface{}, dur time.Duration) error {
	return RDB.Set(ctx, key, value, dur).Err()
}

func RDBExists(key string) bool {
	ok, _ := RDB.Exists(ctx, key).Result()
	return ok == 1
}
