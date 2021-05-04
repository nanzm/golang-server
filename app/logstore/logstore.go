package logstore

import (
	store "dora/app/logstore/core"
	elasticComponent "dora/app/logstore/elastic"
	slsLogComponent "dora/app/logstore/slslog"
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
