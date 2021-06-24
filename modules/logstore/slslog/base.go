package slslogComponent

import (
	"dora/config"
	"dora/modules/logstore/datasource/slslog"
	"dora/pkg/utils"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"google.golang.org/protobuf/proto"
	"time"
)

// sdk 存入阿里云日志服务
func basePutLog(mapData map[string]interface{}) error {
	logs := fmtLog(mapData)
	ins := slslog.GetProducer()
	conf := config.GetSlsLog()
	err := ins.SendLog(conf.Project, conf.LogStore, conf.Topic, conf.Source, logs)
	return err
}

func basePutLogList(mapData []map[string]interface{}) error {
	logs := fmtLogList(mapData)
	ins := slslog.GetProducer()
	conf := config.GetSlsLog()
	err := ins.SendLogList(conf.Project, conf.LogStore, conf.Topic, conf.Source, logs)
	return err
}

// sdk 查询日志数据
func baseQueryLogs(from int64, to int64, queryExp string) (result *sls.GetLogsResponse, err error) {
	ins := slslog.GetClient()
	conf := config.GetSlsLog()
	logs, err := ins.GetLogs(conf.Project, conf.LogStore, conf.Topic, from, to, queryExp, 100, int64(0), true)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func fmtLog(event map[string]interface{}) *sls.Log {
	var content []*sls.LogContent
	for key, value := range event {
		content = append(content, &sls.LogContent{
			Key:   proto.String(key),
			Value: proto.String(utils.SafeJsonMarshal(value)),
		})
	}

	current := uint32(time.Now().Unix())
	return &sls.Log{
		Time:     proto.Uint32(current),
		Contents: content,
	}
}

func fmtLogList(eventList []map[string]interface{}) []*sls.Log {
	list := make([]*sls.Log, 0)

	for _, event := range eventList {
		var content []*sls.LogContent
		for key, value := range event {
			content = append(content, &sls.LogContent{
				Key:   proto.String(key),
				Value: proto.String(utils.SafeJsonMarshal(value)),
			})
		}

		current := uint32(time.Now().Unix())
		list = append(list, &sls.Log{
			Time:     proto.Uint32(current),
			Contents: content,
		})
	}
	return list

}
