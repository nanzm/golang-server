package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

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
