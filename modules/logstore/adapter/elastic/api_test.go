package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
	"testing"
)

func init() {
	logx.Init("./dora.log")
}

func Test_elasticQuery_PvUvTotal(t *testing.T) {
	total, err := NewElasticQuery().PvUvTotal("wdssfar2312312dsad", 1624177745000, 1625041745000)
	if err != nil {
		t.Fatal(err)
	}
	utils.PrettyPrint(total)
}

func Test_elasticQuery_PvUvTrend(t *testing.T) {
	re, err := NewElasticQuery().PvUvTrend("wdssfar2312312dsad", 1624177745000, 1625041745000, 30)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}

func Test_elasticQuery_PagesCount(t *testing.T) {
	re, err := NewElasticQuery().PagesCount("wdssfar2312312dsad", 1624177745000, 1625041745000)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}

func Test_elasticQuery_ErrorCount(t *testing.T) {
	re, err := NewElasticQuery().ErrorCount("wdssfar2312312dsad", 1624177745000, 1625041745000)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}

func Test_elasticQuery_PerfMetricsTrend(t *testing.T) {
	//_, err := NewElasticQuery().PerfMetricsBucket("wdssfar2312312dsad", 1617097374, 1617183785)
	//if err != nil {
	//	print(err)
	//}

	res := ``

	result := &response.PerfMetricsBucket{}

	metrics := gjson.Get(res, "aggregations")
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

	utils.PrettyPrint(result)
}

func Test_elasticQuery_SdkVersionCount(t *testing.T) {
	re, err := NewElasticQuery().SdkVersionCount("wdssfar2312312dsad", 1624177745000, 1625041745000)
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(re)
}