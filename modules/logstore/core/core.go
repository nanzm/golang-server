package store

import "dora/modules/logstore/response"

type Api interface {
	PutData(logData map[string]interface{}) error
	PutListData(listLogData []map[string]interface{}) error

	DefaultQuery(appId string, from, to, interval int64, dataType string) (result interface{}, err error)
	QueryMethods() QueryMethods
}

type QueryMethods interface {
	GetLogByMd5(from, to int64, md5 string) (*response.LogsResponse, error)
	LogCountByMd5(from, to int64, md5 string) (*response.LogCountByMd5Res, error)

	PvUvTotal(appId string, from, to int64) (*response.PvUvTotalRes, error)
	PvUvTrend(appId string, from, to, interval int64) (*response.PvUvTrendRes, error)

	SdkVersionCount(appId string, from, to int64) (*response.SdkVersionCountRes, error)
	CategoryCount(appId string, from, to int64) (*response.CategoryCountRes, error)
	PagesCount(appId string, from, to int64) (*response.PageTotalRes, error)

	// 错误
	ErrorCount(appId string, from, to int64) (*response.ErrorCountRes, error)
	ErrorCountTrend(appId string, from, to, interval int64) (*response.ErrorCountTrendRes, error)

	// 接口错误
	ApiErrorCount(appId string, from, to int64) (*response.ApiErrorCountRes, error)
	ApiErrorTrend(appId string, from, to int64, interval int64) (*response.ApiErrorTrendRes, error)
	ApiErrorList(appId string, from, to int64) (*response.ApiErrorListRes, error)

	// 网络
	PerfNavigationTimingTrend(appId string, from, to int64, interval int64) (*response.PerfNavigationTimingTrendRes, error)
	PerfNavigationTimingValues(appId string, from, to int64) (*response.PerfNavigationTimingValuesRes, error)

	// 资源加载
	PerfDataConsumptionTrend(appId string, from, to int64, interval int64) (*response.PerfDataConsumptionTrendRes, error)
	PerfDataConsumptionValues(appId string, from, to int64) (*response.PerfDataConsumptionValuesRes, error)

	// 性能
	PerfMetricsTrend(appId string, from, to int64, interval int64) (*response.PerfMetricsTrendRes, error)
	PerfMetricsValues(appId string, from, to int64) (*response.PerfMetricsValuesRes, error)

	// 资源加载失败
	ResLoadFailTotalTrend(appId string, from, to, interval int64) (*response.ResLoadFailTotalTrendRes, error)
	ResLoadFailTotal(appId string, from, to int64) (*response.ResLoadFailTotalRes, error)
	ResLoadFailList(appId string, from, to int64) (*response.ResLoadFailListRes, error)

	ProjectIpToCountry(appId string, from, to int64) (*response.ProjectIpToCountryRes, error)
	ProjectIpToProvince(appId string, from, to int64) (*response.ProjectIpToProvinceRes, error)
	ProjectIpToCity(appId string, from, to int64) (*response.ProjectIpToCityRes, error)

	ProjectEventCount(appId string, from, to int64) (*response.ProjectEventCountRes, error)
	ProjectSendMode(appId string, from, to int64) (*response.ProjectSendModeRes, error)
	ProjectVersion(appId string, from, to int64) (*response.ProjectVersionRes, error)
	ProjectUserScreen(appId string, from, to int64) (*response.ProjectUserScreenRes, error)
	ProjectCategory(appId string, from, to int64) (*response.ProjectCategoryRes, error)
	ProjectEnv(appId string, from, to int64) (*response.ProjectEnvRes, error)
}
