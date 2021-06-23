package core

import (
	"dora/modules/api/transit/config"
	"dora/modules/datasource/nsq"
	"dora/modules/datasource/slslog"
	"dora/pkg/utils/logx"
)

func Setup() {
	// log
	conf := config.GetLog()
	logx.Init(conf.File)

	// nsq
	nsq.ProducerInstance()

	// logStore
	slslog.GetProducer()
}

func TearDown() {
	nsq.ProducerTearDown()

	slslog.TearDownProducer()
}
