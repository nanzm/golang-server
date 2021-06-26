package slslog

import (
	"dora/config"
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
	err := GetClient().Close()
	logx.Println("slsLog client closed")
	if err != nil {
		logx.Errorf("%v \n", err)
	}
}
