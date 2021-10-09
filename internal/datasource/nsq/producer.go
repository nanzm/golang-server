package nsq

import (
	"dora/internal/config"
	"dora/pkg/utils/logx"
	"github.com/nsqio/go-nsq"
	"sync"
)

var onceNsq sync.Once
var nsqProducer *nsq.Producer

// 生产
func ProducerInstance() *nsq.Producer {
	onceNsq.Do(func() {
		conf := config.GetNsq()

		// 生产者
		c := nsq.NewConfig()
		p, err := nsq.NewProducer(conf.Address, c)
		if err != nil {
			logx.Panic(err)
		}

		p.SetLogger(&customLog{}, nsq.LogLevelWarning)

		logx.Info("nsq producer ready")
		nsqProducer = p
	})
	return nsqProducer
}


func ProducerTearDown() {
	if nsqProducer != nil {
		logx.Println("stop nsq producer")
		nsqProducer.Stop()
	}

}
