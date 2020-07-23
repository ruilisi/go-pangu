package redis

import (
	"context"
	"fmt"
	"go-jwt/setting"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

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

func Get(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func Set(key string, value interface{}, dur time.Duration) error {
	return RDB.Set(ctx, key, value, dur).Err()
}

func Exists(key string) bool {
	ok, _ := RDB.Exists(ctx, key).Result()
	return ok == 1
}
