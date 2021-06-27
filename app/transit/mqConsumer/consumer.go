package mqConsumer

import (
	"dora/modules/logstore"
	"dora/pkg/utils"
	"dora/pkg/utils/logx"
	"github.com/nsqio/go-nsq"
)

// 消费队列
func Consumer() nsq.Handler {
	return nsq.HandlerFunc(func(message *nsq.Message) error {
		logx.Printf("nsq consumer event： %v", message.Timestamp)
		return msgHandle(message)
	})
}

// 消费队列消息 ————> 放入 日志服务
func msgHandle(message *nsq.Message) error {
	// 解析成 map
	event, err := utils.StringToMapList(message.Body)
	if err != nil {
		logx.Error(err)
		return err
	}

	// 存入日志服务
	client := logstore.GetClient()
	err = client.PutListData(event)
	if err != nil {
		return err
	}

	return nil
}
