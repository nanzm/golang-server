package slslog

import (
	"dora/internal/config"
	"dora/pkg/utils/logx"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"sync"
)

var onceClient sync.Once
var client sls.ClientInterface

func GetClient() sls.ClientInterface {
	onceClient.Do(func() {
		conf := config.GetSlsLog()
		client = initClient(conf)
	})
	return client
}

func initClient(c config.SlsLog) sls.ClientInterface {
	Client := sls.CreateNormalInterface(c.Endpoint, c.AccessKey, c.Secret, "")
	return Client
}

func ClientTearDown() {
	if client != nil {
		err := client.Close()
		if err != nil {
			logx.Errorf("%v \n", err)
		}
		logx.Println("slsLog client closed")
	}
}
