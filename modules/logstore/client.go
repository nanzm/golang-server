package logstore

import (
	"dora/config"
	elasticComponent "dora/modules/logstore/adapter/elastic"
	slsLogComponent "dora/modules/logstore/adapter/slslog"
	"dora/modules/logstore/core"
	"dora/modules/logstore/datasource/slslog"
	"dora/pkg/utils/logx"
)

func GetClient() core.Client {
	conf := config.GetLogStore()
	if conf.Enable == "slsLog" {
		return getSlsClient()
	}
	if conf.Enable == "elastic" {
		return getEsClient()
	}
	logx.Fatalf("%s", "config.yml logStore enable value Invalid! ")
	return nil
}

func TearDown() {
	conf := config.GetLogStore()
	if conf.Enable == "slsLog" {
		slslog.TearDownProducer()
	}
}

// 阿里云 sls
func getSlsClient() core.Client {
	return slsLogComponent.NewSlsLogStore()
}

// elastic
func getEsClient() core.Client {
	return elasticComponent.NewElkLogStore()
}
