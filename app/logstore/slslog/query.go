package slslogComponent

import (
	"dora/app/logstore/core"
	"errors"
	"strconv"
	"strings"
)

type slsLog struct {
}

func NewSlsLogStore() store.Api {
	return &slsLog{}
}

func (s slsLog) PutData(logData map[string]interface{}) error {
	err := basePutLogs(logData)
	return err
}

func (s slsLog) DefaultQuery(appId string, from, to, interval int64, dataType string) (interface{}, error) {
	m := NewSlsQuery()
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
		return m.PerfMetricsTrend(appId, from, to, interval)
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

func (s slsLog) QueryMethods() store.QueryMethods {
	return NewSlsQuery()
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
