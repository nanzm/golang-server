package boot

import (
	"dora/internal/apps/transit/mqConsumer"
	"dora/internal/config"
	"dora/internal/datasource/nsq"
	"dora/internal/initialize"
	"dora/internal/logstore"
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
