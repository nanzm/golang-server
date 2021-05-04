package datasource

import (
	"dora/config"
	"dora/pkg/logger"

	"sync"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	consumerLibrary "github.com/aliyun/aliyun-log-go-sdk/consumer"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
)

type Log struct {
	Conf             config.SlsLog
	Client           sls.ClientInterface
	ProducerInstance *producer.Producer
}

var onceAliLog sync.Once
var slsLog *Log
var consumerWorker *consumerLibrary.ConsumerWorker

func GetSlsInstance() *Log {
	onceAliLog.Do(func() {
		conf := config.GetConf()
		slsLog = &Log{
			Conf:             conf.SlsLog,
			Client:           initClient(conf.SlsLog),
			ProducerInstance: initProducer(conf.SlsLog),
		}
	})
	return slsLog
}

func initProducer(c config.SlsLog) *producer.Producer {
	producerConfig := &producer.ProducerConfig{
		TotalSizeLnBytes:      100 * 1024 * 1024,
		MaxIoWorkerCount:      50,
		MaxBlockSec:           60,
		MaxBatchSize:          512 * 1024,
		LingerMs:              2000,
		Retries:               10,
		MaxReservedAttempts:   11,
		BaseRetryBackoffMs:    100,
		MaxRetryBackoffMs:     50 * 1000,
		AdjustShargHash:       true,
		Buckets:               64,
		MaxBatchCount:         4096,
		NoRetryStatusCodeList: []int{400, 404},
		AllowLogLevel:         "error",
	}

	producerConfig.Endpoint = c.Endpoint
	producerConfig.AccessKeyID = c.AccessKey
	producerConfig.AccessKeySecret = c.Secret
	producerInstance := producer.InitProducer(producerConfig)

	// 启动producer实例
	producerInstance.Start()

	logger.Println("sls producerInstance ready")
	return producerInstance

}

func initClient(c config.SlsLog) sls.ClientInterface {
	Client := sls.CreateNormalInterface(c.Endpoint, c.AccessKey, c.Secret, "")
	return Client
}

func SetUpSlsConsumer(c config.SlsLog, consumerFunc func(int, *sls.LogGroupList) string) {
	option := consumerLibrary.LogHubConfig{
		Endpoint:          c.Endpoint,
		AccessKeyID:       c.AccessKey,
		AccessKeySecret:   c.Secret,
		Project:           c.Project,
		Logstore:          c.LogStore,
		ConsumerGroupName: c.ConsumerGroupName,
		ConsumerName:      c.ConsumerName,
		// This options is used for initialization, will be ignored once consumer group is created and each shard has been started to be consumed.
		// Could be "begin", "end", "specific time format in time stamp", it's log receiving time.
		CursorPosition: consumerLibrary.END_CURSOR,
		IsJsonType:     true,
	}
	consumerWorker = consumerLibrary.InitConsumerWorker(option, consumerFunc)
	consumerWorker.Start()
}

func StopSlsLog() {
	logger.Println("close aliyun sls log producer instance")
	GetSlsInstance().ProducerInstance.SafeClose()

	err := GetSlsInstance().Client.Close()
	logger.Println("close aliyun sls log client")
	if err != nil {
		logger.Errorf("%v \n", err)
	}

	if consumerWorker != nil {
		logger.Println("close aliyun sls log consumer worker")
		consumerWorker.StopAndWait()
	}
}
