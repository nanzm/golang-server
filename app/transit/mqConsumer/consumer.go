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
	event, err := utils.StringToMap(message.Body)
	if err != nil {
		logx.Error(err)
		return err
	}

	//// 添加 md5
	//md5Content := md5AggData(event)
	//
	//// 判断是否创建 issues
	//if md5Content != "" {
	//	event["agg_md5"] = md5Content
	//	if redis.SetExist(constant.Md5ListHas, md5Content) {
	//		// 存在 更新
	//		redis.RedisSetAdd(constant.Md5ListWaitUpdate, md5Content)
	//	} else {
	//		// 不存在 创建
	//		redis.RedisSetAdd(constant.Md5ListWaitCreate, md5Content)
	//	}
	//}

	// 存入日志服务
	client := logstore.GetSlsClient()
	err = client.PutData(event)
	if err != nil {
		return err
	}

	// 测试ES 多存一份日志
	//c2 := logstore.GetEsClient()
	//err2 := c2.PutData(event)
	//if err2 != nil {
	//	logx.Errorf("存入%v \n", err2)
	//}

	return nil
}

func md5AggData(event map[string]interface{}) string {
	if val, ok := event["agg"]; ok {
		s := utils.SafeJsonMarshal(val)
		if s == "" {
			return ""
		} else {
			return utils.Md5([]byte(s))
		}
	}
	return ""
}
