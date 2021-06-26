package boot

import (
	"dora/app/transit/mqConsumer"
	"dora/config"
	"dora/modules/datasource/nsq"
	"dora/modules/logstore"
	"dora/pkg/utils/logx"
)

func Setup() {
	// log
	conf := config.GetTransitLog()
	logx.Init(conf.File)

	// logStore
	logstore.GetClient()

	// nsq
	nsq.ProducerInstance()
	nsq.ConsumerRegister(mqConsumer.Consumer())
}

func TearDown() {
	nsq.ProducerTearDown()
	nsq.ConsumerTearDown()
	logstore.TearDown()
}
