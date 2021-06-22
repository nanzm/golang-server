package datasource

import (
	"dora/config"
	"dora/pkg/utils/logx"

	"sync"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
)

type Log struct {
	Conf             config.SlsLog
	Client           sls.ClientInterface
	ProducerInstance *producer.Producer
}

var onceAliLog sync.Once
var slsLog *Log

func GetSlsInstance() *Log {
	onceAliLog.Do(func() {
		conf := config.GetSlsLog()
		slsLog = &Log{
			Conf:             conf,
			Client:           initClient(conf),
			ProducerInstance: initProducer(conf),
		}
	})
	return slsLog
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

func initClient(c config.SlsLog) sls.ClientInterface {
	Client := sls.CreateNormalInterface(c.Endpoint, c.AccessKey, c.Secret, "")
	return Client
}

func StopSlsLog() {
	logx.Println("close aliyun sls log producer instance")
	GetSlsInstance().ProducerInstance.SafeClose()

	err := GetSlsInstance().Client.Close()
	logx.Println("close aliyun sls log client")
	if err != nil {
		logx.Errorf("%v \n", err)
	}

}
