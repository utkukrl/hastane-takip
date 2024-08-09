// utils/redis_util.go
package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(addr, password string, db int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
}

func SetKey(key string, value interface{}) error {
	return RedisClient.Set(context.Background(), key, value, 0).Err()
}

func GetKey(key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}

func DelKey(key string) error {
	return RedisClient.Del(context.Background(), key).Err()
}
