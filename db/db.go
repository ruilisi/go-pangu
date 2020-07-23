package db

import (
	"context"
	"fmt"
	"time"

	"go-jwt/models"
	"go-jwt/setting"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var RDB *redis.Client
var ctx = context.Background()

func ConnectDB() {
	var err error
	//replace your database
	DB, err = gorm.Open(setting.DatabaseSetting.Type,
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Port,
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Name,
			setting.DatabaseSetting.Password))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password, // no password set
		DB:       setting.RedisSetting.DB,       // use default DB
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
