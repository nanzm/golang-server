package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)


func TestRedisInstance(t *testing.T) {
	const StoreSwitch = "logStoreSwitch"

	result, err := Instance().Get(context.Background(), StoreSwitch).Result()
	if err != nil && err != redis.Nil {
		t.Fatal(err)
		return
	}

	t.Logf("%v \n", result)
}
