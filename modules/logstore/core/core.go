package core

import "dora/modules/logstore/response"

type Client interface {
	PutData(logData map[string]interface{}) error
	PutListData(listLogData []map[string]interface{}) error

	DefaultQuery(appId string, from, to, interval int64, dataType string) (result interface{}, err error)
	QueryMethods() Query
}

type Query interface {
	GetLogByMd5(appId string, from, to int64, md5 string) (*response.LogsResponse, error)
	SearchErrorLog(appId string, from, to int64, searchStr string) (*response.LogList, error)

	PvUvTotal(appId string, from, to int64) (*response.PvUvTotalRes, error)
	PvUvTrend(appId string, from, to, interval int64) (*response.PvUvTrendRes, error)

	ProjectLogCount(appId string) (*response.ProjectLogCountRes, error)
	PagesUrlVisitList(appId string, from, to int64) (*response.PageUrlVisitListRes, error)

	ProjectUserScreenList(appId string, from, to int64) (*response.CountListRes, error)
	ProjectLogTypeList(appId string, from, to int64) (*response.CountListRes, error)
	ProjectEnvList(appId string, from, to int64) (*response.CountListRes, error)
	ProjectVersionList(appId string, from, to int64) (*response.CountListRes, error)
	ProjectSdkVersionList(appId string, from, to int64) (*response.CountListRes, error)

	// 错误
	ErrorCount(appId string, from, to int64) (*response.ErrorCountRes, error)
	ErrorCountTrend(appId string, from, to, interval int64) (*response.CountListRes, error)
	ErrorList(appId string, from, to int64) (*response.ErrorListRes, error)

	// 接口错误
	ApiErrorCount(appId string, from, to int64) (*response.ApiErrorCountRes, error)
	ApiErrorTrend(appId string, from, to int64, interval int64) (*response.CountListRes, error)
	ApiErrorList(appId string, from, to int64) (*response.ApiErrorListRes, error)

	ApiDuration(appId string, from, to int64) (*response.ApiDurationRes, error)
	ApiDurationTrend(appId string, from, to int64, interval int64) (*response.ApiDurationTrendRes, error)
	ApiTopListDuration(appId string, from, to int64) (*response.ApiTopListDurationRes, error)

	// 性能
	PerfMetricsBucket(appId string, from, to int64) (*response.PerfMetricsBucket, error)

	// 资源加载失败
	ResLoadFailTotal(appId string, from, to int64) (*response.ResLoadFailTotalRes, error)
	ResLoadFailTrend(appId string, from, to, interval int64) (*response.CountListRes, error)
	ResLoadFailList(appId string, from, to int64) (*response.CountListRes, error)
	ResTopListDuration(appId string, from, to int64) (*response.ResTopListDurationRes, error)
	ResDuration(appId string, from, to int64) (*response.ResDurationRes, error)
	ResDurationTrend(appId string, from, to, interval int64) (*response.CountListRes, error)
}
