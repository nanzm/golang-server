package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)



const resLoadFailTrend =`{
  "size": 0,
  "query": {
    "bool": {
      "filter": [
        {
          "match": {
            "appId": "<APPID>"
          }
        },
        {
          "match": {
            "type": "resource"
          }
        },
        {
          "range": {
            "ts": {
              "gte": <FORM>,
              "lte": <TO>
            }
          }
        }
      ]
    }
  },
  "aggregations": {
    "trend": {
      "date_histogram": {
        "field": "ts",
        "interval": "<INTERVAL>m",
        "time_zone": "+08:00",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "aggregations": {
        "count": {
          "value_count": {
            "field": "type.keyword"
          }
        },
        "effectUser": {
          "cardinality": {
            "field": "uid.keyword"
          }
        }
      }
    }
  }
}`

const resLoadFailList =`{
  "size": 0,
  "query": {
    "bool": {
      "filter": [
        {
          "match": {
            "appId": "<APPID>"
          }
        },
        {
          "match": {
            "type": "resource"
          }
        },
        {
          "range": {
            "ts": {
              "gte": <FORM>,
              "lte": <TO>
            }
          }
        }
      ]
    }
  },
  "aggregations": {
    "url": {
      "terms": {
        "field": "resource.src.keyword",
        "size": 50,
        "order": {
          "count": "desc"
        }
      },
      "aggregations": {
        "count": {
          "value_count": {
            "field": "type.keyword"
          }
        },
        "effectUser": {
          "cardinality": {
            "field": "uid.keyword"
          }
        }
      }
    }
  }
}`

func (e elasticQuery) ResLoadFailTotalTrend(appId string, from, to, interval int64) (*response.ResLoadFailTotalTrendRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(resLoadFailTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ResLoadFailTotalTrendItemRes, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "doc_count").Num
		eUser := gjson.Get(value.Raw, "uv.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()
		item := &response.ResLoadFailTotalTrendItemRes{
			Count:      int(count),
			EffectUser: int(eUser),
			Ts:         ts,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ResLoadFailTotalTrendRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ResLoadFailTotal(appId string, from, to int64) (*response.ResLoadFailTotalRes, error) {
	panic("implement me")
}

func (e elasticQuery) ResLoadFailList(appId string, from, to int64) (*response.ResLoadFailListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(resLoadFailList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ResLoadFailItemRes, 0)

	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		url := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		effectUser := gjson.Get(value.Raw, "effectUser.value").Num

		item := &response.ResLoadFailItemRes{
			Src:        url,
			Count:      int(count),
			EffectUser: int(effectUser),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ResLoadFailListRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ResDuration(appId string, from, to int64) (*response.ResLoadFailListRes, error) {
	panic("implement me")
}

func (e elasticQuery) ResDurationTrend(appId string, from, to, interval int64) (*response.ResLoadFailListRes, error) {
	panic("implement me")
}
