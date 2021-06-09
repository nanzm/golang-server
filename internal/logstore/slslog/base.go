package slslogComponent

import (
	"dora/internal/datasource"
	"dora/pkg/utils"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"google.golang.org/protobuf/proto"
	"time"
)

// sdk 存入阿里云日志服务
func basePutLogs(mapData map[string]interface{}) error {
	logs := generateLog(uint32(time.Now().Unix()), mapData)
	ins := datasource.GetSlsInstance()
	conf := ins.Conf
	err := ins.ProducerInstance.SendLog(conf.Project, conf.LogStore, conf.Topic, conf.Source, logs)
	return err
}

// sdk 查询日志数据
func baseQueryLogs(from int64, to int64, queryExp string) (result *sls.GetLogsResponse, err error) {
	ins := datasource.GetSlsInstance()
	conf := ins.Conf
	logs, err := ins.Client.GetLogs(conf.Project, conf.LogStore, conf.Topic, from, to, queryExp, 100, int64(0), true)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func generateLog(logTime uint32, addLogMap map[string]interface{}) *sls.Log {
	var content []*sls.LogContent
	for key, value := range addLogMap {
		content = append(content, &sls.LogContent{
			Key:   proto.String(key),
			Value: proto.String(utils.SafeJsonMarshal(value)),
		})
	}
	return &sls.Log{
		Time:     proto.Uint32(logTime),
		Contents: content,
	}
}

//func cleanLogs(logs []*sls.Log) []*sls.Log {
//	result := make([]*sls.Log, 0, len(logs))
//	temp := map[string]struct{}{}
//
//	for _, item := range logs {
//		md5 := getAggMd5Val(item.Contents)
//
//		// 忽略没有 md5
//		if md5 == "" {
//			continue
//		}
//
//		// 去重
//		if _, ok := temp[md5]; !ok {
//			temp[md5] = struct{}{}
//			result = append(result, item)
//		}
//	}
//	return result
//}
//
//func getAggMd5Val(c []*sls.LogContent) string {
//	for _, content := range c {
//		if *content.Key == "agg_md5" {
//			return *content.Value
//		}
//	}
//	return ""
//}
