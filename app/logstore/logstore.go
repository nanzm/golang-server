package logstore

import (
	"context"
	"dora/app/datasource"
	store "dora/app/logstore/core"
	elasticComponent "dora/app/logstore/elastic"
	slsLogComponent "dora/app/logstore/slslog"
)

const StoreSwitch = "logStoreSwitch"

func GetClient() store.Api {
	result, _ := datasource.RedisInstance().Get(context.Background(), StoreSwitch).Result()

	if result == "" {
		return GetSlsClient()
	} else {
		return GetEsClient()
	}
}

func GetEsClient() store.Api {
	return elasticComponent.NewElkLogStore()
}

func GetSlsClient() store.Api {
	return slsLogComponent.NewSlsLogStore()
}
