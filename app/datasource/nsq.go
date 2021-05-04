package datasource

import (
	"dora/config"
	"dora/pkg/logger"

	"sync"

	"github.com/nsqio/go-nsq"
)

var onceNsq sync.Once
var nsqProducer *nsq.Producer
var nsqConsumer *nsq.Consumer

func NsqProducerInstance() *nsq.Producer {
	onceNsq.Do(func() {
		conf := config.GetConf()

		// 生产者
		c := nsq.NewConfig()
		p, err := nsq.NewProducer(conf.Nsq.Address, c)
		if err != nil {
			logger.Panic(err)
		}

		p.SetLogger(&customLog{}, nsq.LogLevelWarning)

		logger.Info("nsq producer ready")
		nsqProducer = p
	})
	return nsqProducer
}

// 消费
func NsqConsumerRegister(conf config.NsqConfig, handler nsq.Handler) {
	con := nsq.NewConfig()
	c, err := nsq.NewConsumer(conf.Topic, conf.Channel, con)
	if err != nil {
		logger.Panic(err)
	}
	c.SetLogger(&customLog{}, nsq.LogLevelWarning)

	c.AddHandler(handler)
	err = c.ConnectToNSQD(conf.Address)
	if err != nil {
		logger.Panic(err)
	}

	logger.Info("nsq consumer ready")
	nsqConsumer = c
}

func StopNsq() {
	logger.Println("stop nsq producer")
	NsqProducerInstance().Stop()

	if nsqConsumer != nil {
		logger.Println("stop nsq consumer")
		nsqConsumer.Stop()
	}
}

type customLog struct {
}

func (c *customLog) Output(_ int, s string) error {
	logger.Warnf("nsq: %v", s)
	return nil
}
