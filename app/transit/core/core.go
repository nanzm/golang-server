package core

import (
	"dora/app/transit/config"
	"dora/app/transit/datasource/nsq"
	"dora/modules/logstore/datasource/slslog"
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
