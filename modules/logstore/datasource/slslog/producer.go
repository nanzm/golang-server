package slslog

import (
	"dora/config"
	"dora/pkg/utils/logx"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"sync"
)

var onceProducer sync.Once
var producerIns *producer.Producer

func GetProducer() *producer.Producer {
	onceProducer.Do(func() {
		conf := config.GetSlsLog()
		producerIns = initProducer(conf)
	})
	return producerIns
}

func initProducer(c config.SlsLog) *producer.Producer {
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.AllowLogLevel = "error"
	producerConfig.Endpoint = c.Endpoint
	producerConfig.AccessKeyID = c.AccessKey
	producerConfig.AccessKeySecret = c.Secret
	producerInstance := producer.InitProducer(producerConfig)

	// 启动producer实例
	producerInstance.Start()

	logx.Println("sls producerInstance ready")
	return producerInstance

}

func TearDownProducer() {
	if producerIns != nil {
		producerIns.SafeClose()
		logx.Println("slsLog producer closed!")
	}
}
