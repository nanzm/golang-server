package elasticComponent

import (
	"dora/config"
	"dora/modules/logstore/core"
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type elasticQuery struct {
	config config.Elastic
}

func NewElasticQuery() core.Api {
	return &elasticQuery{
		config: config.GetElastic(),
	}
}

func (e elasticQuery) GetLogByMd5(appId string, from, to int64, md5 string) (*response.LogsResponse, error) {
	r := strings.NewReplacer(
		core.TplAppId, appId,
		core.TplFrom, strconv.Itoa(int(from)),
		core.TplTo, strconv.Itoa(int(to)),
		core.TplMD5, md5,
	)
	tpl := r.Replace(getLogsByMd5)

	res, err := baseSearch(e.config.Index, tpl)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 转
	count := gjson.Get(string(res), "aggregations.count.value").Num
	effectUser := gjson.Get(string(res), "aggregations.effectUser.value").Num

	l := gjson.Get(string(res), "hits.hits").String()

	var logs []map[string]interface{}
	err = jsoniter.Unmarshal([]byte(l), &logs)
	if err != nil {
		return nil, err
	}

	result := &response.LogsResponse{
		Count:      int(count),
		EffectUser: int(effectUser),
		Logs:       logs,
	}
	return result, nil
}

func (e elasticQuery) LogCountByMd5(appId string, from, to int64, md5 string) (*response.LogCountByMd5Res, error) {
	panic("implement me")
}

func (e elasticQuery) GetErrorList(appId string, from, to int64) (*response.ErrorListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(errorList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ErrorItem, 0)

	buckets := gjson.Get(string(res), "aggregations.md5.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		md5 := gjson.Get(value.Raw, "key").String()
		msg := gjson.Get(value.Raw, "msg.buckets.0.key").String()
		errorStr := gjson.Get(value.Raw, "error.buckets.0.key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		effectUser := gjson.Get(value.Raw, "effectUser.value").Num
		times := gjson.Get(value.Raw, "ts.buckets.#.key").Array()
		first, last := GetFirstAndLastTime(times)

		item := &response.ErrorItem{
			Md5:        md5,
			Msg:        msg,
			Error:      errorStr,
			Count:      int(count),
			EffectUser: int(effectUser),
			FirstAt:    first,
			LastAt:     last,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ErrorListRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) PvUvTotal(appId string, from, to int64) (*response.PvUvTotalRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(pvUvTotal, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	pv := gjson.Get(string(res), "aggregations.pv.value").Num
	uv := gjson.Get(string(res), "aggregations.uv.value").Num

	result := &response.PvUvTotalRes{
		Pv: int(pv),
		Uv: int(uv),
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
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		pv := gjson.Get(value.Raw, "doc_count").Num
		uv := gjson.Get(value.Raw, "uv.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()
		item := &response.PvUvTrendItemRes{
			Pv: int(pv),
			Uv: int(uv),
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
	res, err := baseSearch(e.config.Index, buildQueryTpl(sdkVersionList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.SdkVersionItem, 0)
	buckets := gjson.Get(string(res), "aggregations.sdk.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		version := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "count.value").Num

		item := &response.SdkVersionItem{
			Version: version,
			Count:   int(count),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.SdkVersionCountRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) CategoryCount(appId string, from, to int64) (*response.CategoryCountRes, error) {
	panic("implement me")
}

func (e elasticQuery) PagesCount(appId string, from, to int64) (*response.PageTotalRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(urlPVUv, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.PageTotalItemRes, 0)
	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		url := gjson.Get(value.Raw, "key").String()
		pv := gjson.Get(value.Raw, "pv.value").Num
		uv := gjson.Get(value.Raw, "uv.value").Num

		item := &response.PageTotalItemRes{
			Url: url,
			Pv:  int(pv),
			Uv:  int(uv),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.PageTotalRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ErrorCount(appId string, from, to int64) (*response.ErrorCountRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(errorCount, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	c := gjson.Get(string(res), "aggregations.count.value").Num
	u := gjson.Get(string(res), "aggregations.effectUser.value").Num

	result := &response.ErrorCountRes{
		Count:      int(c),
		EffectUser: int(u),
	}
	return result, nil
}

func (e elasticQuery) ErrorCountTrend(appId string, from, to, interval int64) (*response.ErrorCountTrendRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(errorCountTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ErrorCountTrendItemRes, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "doc_count").Num
		eUser := gjson.Get(value.Raw, "uv.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()

		item := &response.ErrorCountTrendItemRes{
			Count:      int(count),
			EffectUser: int(eUser),
			Ts:         ts,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ErrorCountTrendRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ApiErrorCount(appId string, from, to int64) (*response.ApiErrorCountRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(apiErrorCount, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	c := gjson.Get(string(res), "aggregations.count.value").Num
	u := gjson.Get(string(res), "aggregations.effectUser.value").Num

	result := &response.ApiErrorCountRes{
		Count:      int(c),
		EffectUser: int(u),
	}
	return result, nil
}

func (e elasticQuery) ApiErrorTrend(appId string, from, to int64, interval int64) (*response.ApiErrorTrendRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(apiErrorTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ApiErrorTrendItemRes, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "doc_count").Num
		eUser := gjson.Get(value.Raw, "uv.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()
		item := &response.ApiErrorTrendItemRes{
			Count:      int(count),
			EffectUser: int(eUser),
			Ts:         ts,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ApiErrorTrendRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ApiErrorList(appId string, from, to int64) (*response.ApiErrorListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(apiErrorList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ApiErrorItem, 0)

	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		url := gjson.Get(value.Raw, "key").String()
		method := gjson.Get(value.Raw, "method.buckets.#.key").String()
		et := gjson.Get(value.Raw, "type.buckets.#.key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		effectUser := gjson.Get(value.Raw, "effectUser.value").Num

		item := &response.ApiErrorItem{
			Id:         value.Index,
			Url:        url,
			Method:     method,
			ErrorType:  et,
			Count:      int(count),
			EffectUser: int(effectUser),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ApiErrorListRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) PerfMetricsBucket(appId string, from, to int64) (*response.PerfMetricsBucket, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(performanceBucket, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	result := &response.PerfMetricsBucket{}

	metrics := gjson.Get(string(res), "aggregations")
	metrics.ForEach(func(key, value gjson.Result) bool {

		resultItem := make([]*response.PerfMetricsBucketItem, 0)

		metrics := gjson.Get(value.Raw, "buckets")
		metrics.ForEach(func(i, j gjson.Result) bool {
			name := gjson.Get(j.Raw, "key").String()
			count := gjson.Get(j.Raw, "doc_count").Value()

			resultItem = append(resultItem, &response.PerfMetricsBucketItem{
				Key: name,
				Val: int(count.(float64)),
			})
			return true
		})

		if key.Str == "fp" {
			result.Fp = resultItem
		}
		if key.Str == "fcp" {
			result.Fcp = resultItem
		}
		if key.Str == "lcp" {
			result.Lcp = resultItem
		}
		if key.Str == "fid" {
			result.Fid = resultItem
		}
		if key.Str == "cls" {
			result.Cls = resultItem
		}
		if key.Str == "ttfb" {
			result.Ttfb = resultItem
		}
		return true
	})

	return result, err
}

func (e elasticQuery) PerfXhrTiming(appId string, from, to int64) (*response.PerfDataConsumptionTrendRes, error) {
	panic("implement me")
}

func (e elasticQuery) PerfScriptTiming(appId string, from, to int64) (*response.PerfDataConsumptionTrendRes, error) {
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
