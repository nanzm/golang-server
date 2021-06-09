package datasource

import (
	"dora/config"
	"dora/pkg/logger"

	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var onceRedis sync.Once

func RedisInstance() *redis.Client {
	onceRedis.Do(func() {
		conf := config.GetConf()
		redisClient = newRedisClient(conf.Redis)
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
	logger.Printf("redis ping: %v ", result)
	return client
}

func StopRedisClient() {
	logger.Println("stop redis client")
	err := RedisInstance().Close()
	if err != nil {
		logger.Error(err)
	}
}

func RedisSetAdd(key string, md5 ...interface{}) {
	ctx := context.Background()
	err := RedisInstance().SAdd(ctx, key, md5...).Err()
	if err != nil {
		logger.Errorf("redis set add: %v", err)
	}
}

func RedisSetExist(key string, md5 string) bool {
	ctx := context.Background()
	result, err := RedisInstance().SIsMember(ctx, key, md5).Result()
	if err != nil {
		logger.Errorf("redis set exist: %v", err)
	}
	return result
}
