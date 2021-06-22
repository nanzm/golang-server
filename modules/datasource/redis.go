package datasource

import (
	"dora/config"
	"dora/pkg/utils/logx"

	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var onceRedis sync.Once

func RedisInstance() *redis.Client {
	onceRedis.Do(func() {
		conf := config.GetRedis()
		redisClient = newRedisClient(conf)
	})
	return redisClient
}

func newRedisClient(c config.RedisConfig) *redis.Client {
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

func StopRedisClient() {
	logx.Println("stop redis client")
	err := RedisInstance().Close()
	if err != nil {
		logx.Error(err)
	}
}

func RedisSetAdd(key string, md5 ...interface{}) {
	ctx := context.Background()
	err := RedisInstance().SAdd(ctx, key, md5...).Err()
	if err != nil {
		logx.Errorf("redis set add: %v", err)
	}
}

func RedisSetExist(key string, md5 string) bool {
	ctx := context.Background()
	result, err := RedisInstance().SIsMember(ctx, key, md5).Result()
	if err != nil {
		logx.Errorf("redis set exist: %v", err)
	}
	return result
}
