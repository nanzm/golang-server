package redis

import (
	"dora/internal/config"
	"dora/pkg/utils/logx"

	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var onceRedis sync.Once

func Instance() *redis.Client {
	onceRedis.Do(func() {
		conf := config.GetRedis()
		redisClient = newClient(conf)
	})
	return redisClient
}

func newClient(c config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       c.DB,
	})

	ctx := context.Background()
	result, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	logx.Printf("redis ping: %v ", result)
	return client
}

func StopClient() {
	logx.Println("stop redis client")
	err := Instance().Close()
	if err != nil {
		logx.Error(err)
	}
}
