package datasource

import (
	"context"
	"dora/config"
	"github.com/go-redis/redis/v8"
	"testing"
)

func init() {
	config.ParseConf("../../config.yml")
}

func TestRedisInstance(t *testing.T) {
	const StoreSwitch = "logStoreSwitch"

	result, err := RedisInstance().Get(context.Background(), StoreSwitch).Result()
	if err != nil && err != redis.Nil {
		t.Fatal(err)
		return
	}

	t.Logf("%v \n", result)
}
