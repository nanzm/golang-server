package logstore

import (
	elasticComponent "dora/internal/logstore/adapter/elastic"
	"dora/internal/logstore/core"
)

func GetClient() core.Client {
	return getEsClient()
}

func TearDown() {
}

// 阿里云 sls
//func getSlsClient() core.Client {
//	return slsLogComponent.NewSlsLogStore()
//}

// elastic
func getEsClient() core.Client {
	return elasticComponent.NewElkLogStore()
}
