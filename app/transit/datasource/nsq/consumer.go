package nsq

import (
	"dora/app/transit/config"
	"dora/pkg/utils/logx"
	"github.com/nsqio/go-nsq"
)

var nsqConsumer *nsq.Consumer

// 消费
func ConsumerRegister(conf config.NsqConfig, handler nsq.Handler) {
	con := nsq.NewConfig()
	c, err := nsq.NewConsumer(conf.Topic, conf.Channel, con)
	if err != nil {
		logx.Panic(err)
	}
	c.SetLogger(&customLog{}, nsq.LogLevelError)

	c.AddHandler(handler)
	err = c.ConnectToNSQD(conf.Address)
	if err != nil {
		logx.Panic(err)
	}

	logx.Info("nsq consumer ready")
	nsqConsumer = c
}

func ConsumerTearDown() {
	if nsqConsumer != nil {
		logx.Println("stop nsq consumer")
		nsqConsumer.Stop()
	}
}