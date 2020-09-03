package auth

import (
	"context"
	"errors"
	redis "github.com/go-redis/redis/v8"
	"github.com/mohammed-maher/fastapi/config"
	"log"
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

func Del(key string) error {
	log.Println("called ",key)
	if deleted, err := RedisClient.Del(ctx, key).Result(); err != nil || deleted==0{
		return errors.New("key could not be deleted")
	}
	return nil
}

