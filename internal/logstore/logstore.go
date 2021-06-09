package logstore

import (
	store "dora/internal/logstore/core"
	elasticComponent "dora/internal/logstore/elastic"
	slsLogComponent "dora/internal/logstore/slslog"
)

func GetClient() store.Api {
	//result, _ := datasource.RedisInstance().Get(context.Background(), constant.RedisKeyStoreSwitch).Result()
	//if result == "" {
	//	return GetSlsClient()
	//} else {
	//	return GetEsClient()
	//}

	return GetSlsClient()
}

func GetEsClient() store.Api {
	return elasticComponent.NewElkLogStore()
}

func GetSlsClient() store.Api {
	return slsLogComponent.NewSlsLogStore()
}
