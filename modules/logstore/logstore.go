package logstore

import (
	store "dora/modules/logstore/core"
	elasticComponent "dora/modules/logstore/elastic"
	slsLogComponent "dora/modules/logstore/slslog"
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
