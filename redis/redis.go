package redis

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RDB *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	u, err := url.Parse(viper.Get("REDIS_URL").(string))
	if err != nil {
		panic(err.Error())
	}

	db, err := strconv.Atoi(u.Path[1:])
	if err != nil {
		panic("Redis url format error")
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: "", // no password set
		DB:       db, // use default DB
	})

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
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
