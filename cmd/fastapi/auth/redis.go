package auth

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"github.com/mohammed-maher/fastapi/config"
	"time"
)

var RedisClient *redis.Client
var ctx=context.Background()

func SetupRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.DSN,
		Password: config.Config.Redis.Password,
	})
}

func Set(key string,val interface{},exp time.Duration) error{
	return RedisClient.Set(ctx,key,val,exp).Err()
}

func Get(key string) *redis.StringCmd{
	return RedisClient.Get(ctx,key)
}
