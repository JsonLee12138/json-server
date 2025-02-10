package core

import (
	"context"
	"fmt"

	"github.com/JsonLee12138/json-server/pkg/configs"
	"github.com/go-redis/redis/v8"
)

func NewRedis(cnf configs.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cnf.Addr(),
		Password: cnf.Password,
		DB:       cnf.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	if Mode() == DevMode {
		fmt.Println("redis连接成功:", pong)
	}
	return client
}
