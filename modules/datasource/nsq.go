package datasource

import (
	"dora/config"
	"dora/pkg/utils/logx"

	"sync"

	"github.com/nsqio/go-nsq"
)

var onceNsq sync.Once
var nsqProducer *nsq.Producer
var nsqConsumer *nsq.Consumer

// 生产
func NsqProducerInstance() *nsq.Producer {
	onceNsq.Do(func() {
		conf := config.GetNsq()

		// 生产者
		c := nsq.NewConfig()
		p, err := nsq.NewProducer(conf.Address, c)
		if err != nil {
			logx.Panic(err)
		}

		p.SetLogger(&customLog{}, nsq.LogLevelError)

		logx.Info("nsq producer ready")
		nsqProducer = p
	})
	return nsqProducer
}

// 消费
func NsqConsumerRegister(conf config.NsqConfig, handler nsq.Handler) {
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

type customLog struct {
}

func (c *customLog) Output(_ int, s string) error {
	logx.Warnf("nsq: %v", s)
	return nil
}

func StopNsq() {
	if nsqProducer != nil {
		logx.Println("stop nsq producer")
		nsqProducer.Stop()
	}

	if nsqConsumer != nil {
		logx.Println("stop nsq consumer")
		nsqConsumer.Stop()
	}
}
