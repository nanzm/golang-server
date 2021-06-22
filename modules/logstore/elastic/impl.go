package elasticComponent

import (
	"dora/config"
	store "dora/modules/logstore/core"
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

type elasticQuery struct {
	config config.Elastic
}

func (e elasticQuery) GetLogByMd5(from, to int64, md5 string) (*response.LogsResponse, error) {
	panic("implement me")
}

func (e elasticQuery) LogCountByMd5(from, to int64, md5 string) (*response.LogCountByMd5Res, error) {
	panic("implement me")
}

func (e elasticQuery) PvUvTotal(appId string, from, to int64) (*response.PvUvTotalRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(pvUvTotal, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	pv := gjson.Get(string(res), "aggregations.pv.value").Value()
	uv := gjson.Get(string(res), "aggregations.uv.value").Value()

	result := &response.PvUvTotalRes{
		Pv: int(pv.(float64)),
		Uv: int(uv.(float64)),
	}
	return result, nil
}

func (e elasticQuery) PvUvTrend(appId string, from, to, interval int64) (*response.PvUvTrendRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(pvUvTotalTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.PvUvTrendItemRes, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.pvTrend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		// {"key_as_string":"2021-03-30 09:30:00","key":1617096600000,"doc_count":4,"uv":{"value":1}}
		pv := gjson.Get(value.Raw, "doc_count").Value()
		uv := gjson.Get(value.Raw, "uv.value").Value()
		ts := gjson.Get(value.Raw, "key_as_string").String()
		item := &response.PvUvTrendItemRes{
			Pv: int(pv.(float64)),
			Uv: int(uv.(float64)),
			Ts: ts,
		}
		logs = append(logs, item)
		return true // keep iterating
	})

	result := &response.PvUvTrendRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) SdkVersionCount(appId string, from, to int64) (*response.SdkVersionCountRes, error) {
	panic("implement me")
}

func (e elasticQuery) CategoryCount(appId string, from, to int64) (*response.CategoryCountRes, error) {
	panic("implement me")
}

func (e elasticQuery) PagesCount(appId string, from, to int64) (*response.PageTotalRes, error) {
	panic("implement me")
}

func (e elasticQuery) ErrorCount(appId string, from, to int64) (*response.ErrorCountRes, error) {
	panic("implement me")
}

func (e elasticQuery) ErrorCountTrend(appId string, from, to, interval int64) (*response.ErrorCountTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) ApiErrorCount(appId string, from, to int64) (*response.ApiErrorCountRes, error) {
	panic("implement me")
}

func (e elasticQuery) ApiErrorTrend(appId string, from, to int64, interval int64) (*response.ApiErrorTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) ApiErrorList(appId string, from, to int64) (*response.ApiErrorListRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfNavigationTimingTrend(appId string, from, to int64, interval int64) (*response.PerfNavigationTimingTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfNavigationTimingValues(appId string, from, to int64) (*response.PerfNavigationTimingValuesRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfDataConsumptionTrend(appId string, from, to int64, interval int64) (*response.PerfDataConsumptionTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfDataConsumptionValues(appId string, from, to int64) (*response.PerfDataConsumptionValuesRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfMetricsTrend(appId string, from, to int64, interval int64) (*response.PerfMetricsTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfMetricsValues(appId string, from, to int64) (*response.PerfMetricsValuesRes, error) {
	panic("implement me")
}

func (e elasticQuery) ResLoadFailTotalTrend(appId string, from, to, interval int64) (*response.ResLoadFailTotalTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) ResLoadFailTotal(appId string, from, to int64) (*response.ResLoadFailTotalRes, error) {
	panic("implement me")
}

func (e elasticQuery) ResLoadFailList(appId string, from, to int64) (*response.ResLoadFailListRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectIpToCountry(appId string, from, to int64) (*response.ProjectIpToCountryRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectIpToProvince(appId string, from, to int64) (*response.ProjectIpToProvinceRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectIpToCity(appId string, from, to int64) (*response.ProjectIpToCityRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectEventCount(appId string, from, to int64) (*response.ProjectEventCountRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectSendMode(appId string, from, to int64) (*response.ProjectSendModeRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectVersion(appId string, from, to int64) (*response.ProjectVersionRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectUserScreen(appId string, from, to int64) (*response.ProjectUserScreenRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectCategory(appId string, from, to int64) (*response.ProjectCategoryRes, error) {
	panic("implement me")
}

func (e elasticQuery) ProjectEnv(appId string, from, to int64) (*response.ProjectEnvRes, error) {
	panic("implement me")
}

func NewElasticQuery() store.QueryMethods {
	return &elasticQuery{
		config: config.GetElastic(),
	}
}

//
//func (e elasticQuery) GetLogByMd5(from, to int64, md5 string) (*store.LogsResponse, error) {
//	return nil, errors.New("be under construction")
//}
//
//func (e elasticQuery) CountLogByMd5(from, to int64, md5 string) (*store.LogsResponse, error) {
//	return nil, errors.New("be under construction")
//}
//
//func (e elasticQuery) PvUvTotal(appId string, from, to int64) (*store.LogsResponse, error) {
//	res, err := baseSearch(e.config.Index, buildQueryTpl(pvUvTotal, appId, from, to))
//	if err != nil {
//		logx.Error(err)
//		return nil, err
//	}
//
//	pv := gjson.Get(string(res), "aggregations.pv.value")
//	uv := gjson.Get(string(res), "aggregations.uv.value")
//
//	result := &store.LogsResponse{
//		Count: 1,
//		Logs:  []map[string]string{{"total": pv.String(), "user": uv.String()}},
//	}
//	return result, nil
//}
//
//func (e elasticQuery) PvUvTrend(appId string, from, to, interval int64) (*store.LogsResponse, error) {
//	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(pvUvTotalTrend, appId, from, to, interval))
//	if err != nil {
//		logx.Error(err)
//		return nil, err
//	}
//
//	logs := make([]map[string]string, 0)
//
//	// 遍历
//	buckets := gjson.Get(string(res), "aggregations.pvTrend.buckets")
//	buckets.ForEach(func(key, value gjson.Result) bool {
//		item := make(map[string]string)
//		// {"key_as_string":"2021-03-30 09:30:00","key":1617096600000,"doc_count":4,"uv":{"value":1}}
//		item["ts"] = gjson.Get(value.Raw, "key_as_string").String()
//		item["pv"] = gjson.Get(value.Raw, "doc_count").String()
//		item["uv"] = gjson.Get(value.Raw, "uv.value").String()
//		logs = append(logs, item)
//		return true // keep iterating
//	})
//
//	result := &store.LogsResponse{
//		Count: int64(len(logs)),
//		Logs:  logs,
//	}
//	return result, nil
//}
//
//func (e elasticQuery) SdkVersionCount(appId string, from, to int64) (*store.LogsResponse, error) {
//	return nil, errors.New("be under construction")
//}
//
//func (e elasticQuery) CategoryCount(appId string, from, to int64) (*store.LogsResponse, error) {
//	return nil, errors.New("be under construction")
//}
//
//func (e elasticQuery) EntryPage(appId string, from, to int64) (*store.LogsResponse, error) {
//	res, err := baseSearch(e.config.Index, buildQueryTpl(entryPage, appId, from, to))
//	if err != nil {
//		logx.Error(err)
//		return nil, err
//	}
//
//	logs := make([]map[string]string, 0)
//
//	// 遍历
//	buckets := gjson.Get(string(res), "aggregations.entryPage.buckets")
//	buckets.ForEach(func(key, value gjson.Result) bool {
//		item := make(map[string]string)
//		// {"key_as_string":"2021-03-30 09:30:00","key":1617096600000,"doc_count":4,"uv":{"value":1}}
//		item["url"] = gjson.Get(value.Raw, "key").String()
//		item["c"] = gjson.Get(value.Raw, "pv.value").String()
//		item["u"] = gjson.Get(value.Raw, "uv.value").String()
//		logs = append(logs, item)
//		return true // keep iterating
//	})
//
//	result := &store.LogsResponse{
//		Count: int64(len(logs)),
//		Logs:  logs,
//	}
//	return result, nil
//}
//
//func (e elasticQuery) ErrorCount(appId string, from, to int64) (*store.LogsResponse, error) {
//	res, err := baseSearch(e.config.Index, buildQueryTpl(errorCount, appId, from, to))
//	if err != nil {
//		logx.Error(err)
//		return nil, err
//	}
//
//	pv := gjson.Get(string(res), "aggregations.pv.value")
//	uv := gjson.Get(string(res), "aggregations.uv.value")
//
//	result := &store.LogsResponse{
//		Count: 1,
//		Logs:  []map[string]string{{"total": pv.String(), "user": uv.String()}},
//	}
//	return result, nil
//}
