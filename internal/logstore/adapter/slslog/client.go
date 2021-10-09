package slslogComponent

import (
	"dora/internal/config"
	"dora/internal/datasource/slslog"
	"dora/internal/logstore/core"
	"dora/pkg/utils"
	"errors"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
	"time"
)

type slsLog struct {
	config   config.SlsLog
	producer *producer.Producer
	client   sls.ClientInterface
}
//
//func NewSlsLogStore() core.Client {
//	return &slsLog{
//	}
//}

func (s slsLog) PutData(logItem map[string]interface{}) error {
	logs := fmtLog(logItem)
	ins := slslog.GetProducer()
	conf := config.GetSlsLog()
	err := ins.SendLog(conf.Project, conf.LogStore, conf.Topic, conf.Source, logs)
	return err
}

func (s slsLog) PutListData(logList []map[string]interface{}) error {
	logs := fmtLogList(logList)
	ins := slslog.GetProducer()
	conf := config.GetSlsLog()
	err := ins.SendLogList(conf.Project, conf.LogStore, conf.Topic, conf.Source, logs)
	return err
}

func (s slsLog) DefaultQuery(appId string, from, to, interval int64, dataType string) (interface{}, error) {
	//m := NewSlsQuery()
	//switch dataType {
	//case "pvUvTotal":
	//	return m.PvUvTotal(appId, from, to)
	//case "pvUvTrend":
	//	return m.PvUvTrend(appId, from, to, interval)
	//case "sdkVersionCount":
	//	return m.SdkVersionCount(appId, from, to)
	//case "categoryCount":
	//	return m.CategoryCount(appId, from, to)
	//case "entryPage":
	//	return m.PagesCount(appId, from, to)
	//
	//// 错误
	//case "errorCount":
	//	return m.ErrorCount(appId, from, to)
	//case "errorCountTrend":
	//	return m.ErrorCountTrend(appId, from, to, interval)
	//case "apiErrorCount":
	//	return m.ApiErrorCount(appId, from, to)
	//case "apiErrorTrend":
	//	return m.ApiErrorTrend(appId, from, to, interval)
	//case "apiErrorList":
	//	return m.ApiErrorList(appId, from, to)
	//
	//// 资源加载错误
	//case "resLoadFailTotalTrend":
	//	return m.ResLoadFailTotalTrend(appId, from, to, interval)
	//case "resLoadFailTotal":
	//	return m.ResLoadFailTotal(appId, from, to)
	//case "resLoadFailList":
	//	return m.ResLoadFailList(appId, from, to)
	//
	//// 性能
	//case "perfMetrics":
	//	return m.PerfMetricsBucket(appId, from, to)
	//
	//case "projectEventCount":
	//	return m.ProjectEventCount(appId, from, to)
	//case "projectSendMode":
	//	return m.ProjectSendMode(appId, from, to)
	//case "projectEnv":
	//	return m.ProjectEnv(appId, from, to)
	//case "projectVersion":
	//	return m.ProjectVersion(appId, from, to)
	//case "projectUserScreen":
	//	return m.ProjectUserScreen(appId, from, to)
	//case "projectCategory":
	//	return m.ProjectCategory(appId, from, to)
	//}

	return nil, errors.New("暂无该指标")
}

func (s slsLog) QueryMethods() core.Client {
	//return NewSlsQuery()
	return nil
}

func buildQueryExp(appId string, queryTpl string) (tpl string, err error) {
	r := strings.NewReplacer("fca5deec-a9db-4dac-a4db-b0f4610d16a5", appId)
	// 替换
	res := r.Replace(queryTpl)
	return res, nil
}

func buildQueryTrendExp(appId string, interval int64, queryTpl string) (tpl string, err error) {
	r := strings.NewReplacer(
		"fca5deec-a9db-4dac-a4db-b0f4610d16a5", appId,
		"2h", strconv.FormatInt(interval, 10)+"m",
	)
	// 替换
	res := r.Replace(queryTpl)
	return res, nil
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
