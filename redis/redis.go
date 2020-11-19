package redis

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"go-pangu/conf"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func newRDB() *redis.Client {
	u, err := url.Parse(conf.GetEnv("REDIS_URL"))
	if err != nil {
		panic(err.Error())
	}

	password, _ := u.User.Password()
	db, err := strconv.Atoi(u.Path[1:])
	if err != nil {
		panic("Redis url format error")
	}

	return redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

//连接redis数据库
func ConnectRedis() {
	RDB = newRDB()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func getRDB() *redis.Client {
	if RDB != nil {
		return RDB
	}
	return newRDB()
}

//redis操作 函数命名与redis实际操作中一致
func Get(key string) string {
	value, err := getRDB().Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func Set(key string, value interface{}) {
	_, err := getRDB().Set(ctx, key, value, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func SetEx(key string, value interface{}, dur time.Duration) {
	_, err := getRDB().Set(ctx, key, value, dur).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func SetNx(key string, value interface{}, dur time.Duration) bool {
	ok, err := getRDB().SetNX(ctx, key, value, dur).Result()
	if err != nil {
		fmt.Println(err)
	}
	return ok
}

func Del(key string) {
	_, err := getRDB().Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func Expire(key string, dur time.Duration) {
	_, err := getRDB().Expire(ctx, key, dur).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func Exists(key string) bool {
	ok, err := getRDB().Exists(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return ok == 1
}

func HSet(key string, values ...interface{}) {
	_, err := getRDB().HSet(ctx, key, values).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func HGetAll(key string) map[string]string {
	value, err := getRDB().HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func HGet(key, field string) string {
	value, err := getRDB().HGet(ctx, key, field).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func HDel(key, field string) {
	_, err := getRDB().HDel(ctx, key, field).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func Smembers(key string) []string {
	values, err := getRDB().SMembers(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return values
}

func Keys(key string) []string {
	values, err := getRDB().Keys(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return values
}

func SPop(key string) string {
	value, err := getRDB().SPop(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func SAdd(key string, members ...interface{}) {
	_, err := getRDB().SAdd(ctx, key, members).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func HExists(key, field string) bool {
	value, err := getRDB().HExists(ctx, key, field).Result()
	if err != nil {
		fmt.Println(err)
	}
	return value
}

func IncrBy(key string, value int64) int64 {
	ret, err := getRDB().IncrBy(ctx, key, value).Result()
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func Do(cmd ...interface{}) (interface{}, error) {
	return getRDB().Do(ctx, cmd...).Result()
}
