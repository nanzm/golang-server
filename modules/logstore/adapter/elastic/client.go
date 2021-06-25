package elasticComponent

import (
	"bytes"
	"dora/config"
	"dora/modules/logstore/core"
	"dora/modules/logstore/datasource/elastic"
	"dora/pkg/utils/logx"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"strconv"
	"strings"
)

type elkLog struct {
	config config.Elastic
	client *elasticsearch.Client
}

func NewElkLogStore() core.Client {
	return &elkLog{
		config: config.GetElastic(),
		client: elastic.GetClient(),
	}
}

func (e elkLog) PutData(logData map[string]interface{}) error {
	byteLogs, err := jsoniter.Marshal(logData)
	if err != nil {
		return err
	}

	es := elastic.GetClient()
	result, err := es.Index(
		e.config.Index,
		bytes.NewReader(byteLogs),
		es.Index.WithRefresh("true"),
		es.Index.WithPretty(),
	)
	if err != nil {
		logx.Printf("%v", err)
		return err
	}

	if result.StatusCode >= 400 {
		logx.Warnf("%v", result)
	}

	return nil
}

func (e elkLog) PutListData(logList []map[string]interface{}) error {
	// todo bulk 批量插入api
	for _, log := range logList {
		err := e.PutData(log)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e elkLog) DefaultQuery(appId string, from, to, interval int64, dataType string) (interface{}, error) {
	m := NewElasticQuery()
	switch dataType {
	case "pvUvTotal":
		return m.PvUvTotal(appId, from, to)
	case "pvUvTrend":
		return m.PvUvTrend(appId, from, to, interval)
	case "sdkVersionCount":
		return m.SdkVersionCount(appId, from, to)
	case "categoryCount":
		return m.CategoryCount(appId, from, to)
	case "entryPage":
		return m.PagesCount(appId, from, to)

	// 错误
	case "errorCount":
		return m.ErrorCount(appId, from, to)
	case "errorCountTrend":
		return m.ErrorCountTrend(appId, from, to, interval)
	case "apiErrorCount":
		return m.ApiErrorCount(appId, from, to)
	case "apiErrorTrend":
		return m.ApiErrorTrend(appId, from, to, interval)
	case "apiErrorList":
		return m.ApiErrorList(appId, from, to)

	// 资源加载错误
	case "resLoadFailTotalTrend":
		return m.ResLoadFailTotalTrend(appId, from, to, interval)
	case "resLoadFailTotal":
		return m.ResLoadFailTotal(appId, from, to)
	case "resLoadFailList":
		return m.ResLoadFailList(appId, from, to)

	// 性能
	case "perfNavigationTiming":
		return m.PerfNavigationTimingTrend(appId, from, to, interval)
	case "perfNavigationTimingValues":
		return m.PerfNavigationTimingValues(appId, from, to)
	case "perfDataConsumption":
		return m.PerfDataConsumptionTrend(appId, from, to, interval)
	case "perfDataConsumptionValues":
		return m.PerfDataConsumptionValues(appId, from, to)
	case "perfMetrics":
		return m.PerfMetricsBucket(appId, from, to)
	case "perfMetricsValues":
		return m.PerfMetricsValues(appId, from, to)

	case "projectEventCount":
		return m.ProjectEventCount(appId, from, to)
	case "projectSendMode":
		return m.ProjectSendMode(appId, from, to)
	case "projectEnv":
		return m.ProjectEnv(appId, from, to)
	case "projectVersion":
		return m.ProjectVersion(appId, from, to)
	case "projectUserScreen":
		return m.ProjectUserScreen(appId, from, to)
	case "projectCategory":
		return m.ProjectCategory(appId, from, to)
	}

	return nil, errors.New("暂无该指标")
}

func (e elkLog) QueryMethods() core.Api {
	return NewElasticQuery()
	//return nil
}

func buildQueryTpl(tpl string, appId string, from, to int64) string {
	r := strings.NewReplacer(
		"fca5deec-a9db-4dac-a4db-b0f4610d16a5", appId,
		"<FORM>", strconv.Itoa(int(from)),
		"<TO>", strconv.Itoa(int(to)),
	)
	// 替换
	res := r.Replace(tpl)
	return res
}

func buildQueryTrendTpl(tpl string, appId string, from, to, interval int64) string {
	r := strings.NewReplacer(
		"fca5deec-a9db-4dac-a4db-b0f4610d16a5", appId,
		"<FORM>", strconv.Itoa(int(from)),
		"<TO>", strconv.Itoa(int(to)),
		"<INTERVAL>", strconv.Itoa(int(interval)),
	)
	// 替换
	res := r.Replace(tpl)
	return res
}

func baseSearch(Index string, queryTpl string) ([]byte, error) {
	//fmt.Printf("%v \n", queryTpl)
	es := elastic.GetClient()

	fmt.Println("------------------------------------")
	res, err := es.Search(
		es.Search.WithIndex(Index),
		es.Search.WithBody(strings.NewReader(queryTpl)),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	// r.StatusCode > 299
	if res.IsError() {
		return nil, errors.New(fmt.Sprintf("%v", res))
	}

	s, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}
