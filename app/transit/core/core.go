package core

import (
	"dora/app/transit/datasource/nsq"
	"dora/app/transit/mqConsumer"
	"dora/config"
	"dora/modules/logstore/datasource/slslog"
	"dora/pkg/utils/logx"
)

func Setup() {
	// log
	conf := config.GetTransitLog()
	logx.Init(conf.File)

	// logStore
	slslog.GetProducer()

	// nsq
	nsq.ProducerInstance()
	nsq.ConsumerRegister(mqConsumer.Consumer())
}

func TearDown() {
	nsq.ProducerTearDown()
	nsq.ConsumerTearDown()

	slslog.TearDownProducer()
}
