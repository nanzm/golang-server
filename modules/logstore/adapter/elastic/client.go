package elasticComponent

import (
	"bytes"
	"context"
	"dora/config"
	"dora/modules/datasource/elastic"
	"dora/modules/logstore/core"
	"dora/pkg/utils/logx"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"time"
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

// 批量插入
func (e elkLog) PutListData(logList []map[string]interface{}) error {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         e.config.Index,   // The default index name
		Client:        e.client,         // The Elasticsearch client
		NumWorkers:    runtime.NumCPU(), // The number of worker goroutines
		FlushBytes:    int(5e+6),        // The flush threshold in bytes
		FlushInterval: 30 * time.Second, // The periodic flush interval
	})

	if err != nil {
		logx.Errorf("Error bulk creating the indexer: %s", err)
		return err
	}

	for _, log := range logList {
		// to byte
		data, err := jsoniter.Marshal(log)
		if err != nil {
			logx.Errorf("Cannot encode log: %s", err)
			return err
		}
		// add item to the BulkIndexer
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "create",
				Body:   bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						logx.Errorf("BulkIndexerItem OnFailure ERROR: %s", err)
					} else {
						logx.Errorf("BulkIndexerItem OnFailure ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			logx.Errorf("bi.Add error: %s", err)
			return err
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		logx.Errorf("Unexpected close error: %s", err)
		return err
	}

	biStats := bi.Stats()
	s, _ := jsoniter.Marshal(biStats)
	logx.Infof("bulk index item stats: %v", string(s))

	return nil
}

func (e elkLog) DefaultQuery(appId string, from, to, interval int64, dataType string) (interface{}, error) {
	m := NewElasticQuery()
	switch dataType {
	case core.PvUvTotal:
		return m.PvUvTotal(appId, from, to)
	case core.PvUvTrend:
		return m.PvUvTrend(appId, from, to, interval)

	// 错误
	case core.ErrorCount:
		return m.ErrorCount(appId, from, to)
	case core.ErrorCountTrend:
		return m.ErrorCountTrend(appId, from, to, interval)

	//	api
	case core.ApiErrorCount:
		return m.ApiErrorCount(appId, from, to)
	case core.ApiErrorTrend:
		return m.ApiErrorTrend(appId, from, to, interval)
	case core.ApiErrorList:
		return m.ApiErrorList(appId, from, to)
	case core.ApiDuration:
		return m.ApiDuration(appId, from, to)
	case core.ApiDurationTrend:
		return m.ApiDurationTrend(appId, from, to, interval)

	// 资源加载错误
	case core.ResLoadFailTotalTrend:
		return m.ResLoadFailTotalTrend(appId, from, to, interval)
	case core.ResLoadFailTotal:
		return m.ResLoadFailTotal(appId, from, to)
	case core.ResLoadFailList:
		return m.ResLoadFailList(appId, from, to)
	case core.ResDuration:
		return m.ResDuration(appId, from, to)
	case core.ResDurationTrend:
		return m.ResDurationTrend(appId, from, to, interval)

	// 性能
	case core.PerfMetrics:
		return m.PerfMetricsBucket(appId, from, to)

	case core.SdkVersionCount:
		return m.SdkVersionCount(appId, from, to)
	case core.CategoryCount:
		return m.CategoryCount(appId, from, to)
	case core.EntryPage:
		return m.PagesCount(appId, from, to)

	case core.ProjectEventCount:
		return m.ProjectEventCount(appId, from, to)
	case core.ProjectSendMode:
		return m.ProjectSendMode(appId, from, to)
	case core.ProjectEnv:
		return m.ProjectEnv(appId, from, to)
	case core.ProjectVersion:
		return m.ProjectVersion(appId, from, to)
	case core.ProjectUserScreen:
		return m.ProjectUserScreen(appId, from, to)
	case core.ProjectCategory:
		return m.ProjectCategory(appId, from, to)
	}

	return nil, errors.New("暂无该指标")
}

func (e elkLog) QueryMethods() core.Query {
	return NewElasticQuery()
}

func buildQueryTpl(tpl string, appId string, from, to int64) string {
	r := strings.NewReplacer(
		core.TplAppId, appId,
		core.TplFrom, strconv.Itoa(int(from)),
		core.TplTo, strconv.Itoa(int(to)),
	)
	// 替换
	res := r.Replace(tpl)
	return res
}

func buildQueryTrendTpl(tpl string, appId string, from, to, interval int64) string {
	r := strings.NewReplacer(
		core.TplAppId, appId,
		core.TplFrom, strconv.Itoa(int(from)),
		core.TplTo, strconv.Itoa(int(to)),
		core.TplInterval, strconv.Itoa(int(interval)),
	)
	// 替换
	res := r.Replace(tpl)
	return res
}

func baseSearch(Index string, queryTpl string) ([]byte, error) {
	fmt.Printf("%v \n", queryTpl)
	es := elastic.GetClient()

	res, err := es.Search(
		es.Search.WithIndex(Index),
		es.Search.WithBody(strings.NewReader(queryTpl)),
		es.Search.WithPretty(),
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
