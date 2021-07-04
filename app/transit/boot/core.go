package boot

import (
	"dora/app/transit/mqConsumer"
	"dora/config"
	"dora/modules/datasource/nsq"
	"dora/modules/initialize"
	"dora/modules/logstore"
	"dora/pkg/utils/logx"
)

func Setup() {
	// log
	conf := config.GetTransitLog()
	logx.Init(conf.File)

	// logStore
	logstore.GetClient()

	// use mapping create index
	initialize.InitElasticIndex()

	// nsq
	nsq.ProducerInstance()
	nsq.ConsumerRegister(mqConsumer.Consumer())
}

func TearDown() {
	nsq.ProducerTearDown()
	nsq.ConsumerTearDown()
	logstore.TearDown()
}
