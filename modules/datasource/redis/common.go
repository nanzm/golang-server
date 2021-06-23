package redis

import (
	"context"
	"dora/pkg/utils/logx"
)

func SetAdd(key string, md5 ...interface{}) {
	ctx := context.Background()
	err := Instance().SAdd(ctx, key, md5...).Err()
	if err != nil {
		logx.Errorf("redis set add: %v", err)
	}
}

func SetExist(key string, md5 string) bool {
	ctx := context.Background()
	result, err := Instance().SIsMember(ctx, key, md5).Result()
	if err != nil {
		logx.Errorf("redis set exist: %v", err)
	}
	return result
}
