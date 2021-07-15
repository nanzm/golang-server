package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

const errorCount = `{
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
            "type": "error"
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
    "count": {
      "cardinality": {
        "field": "eid.keyword"
      }
    },
    "effectUser": {
      "cardinality": {
        "field": "uid.keyword"
      }
    }
  }
}`

const errorCountTrend = `{
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
            "type": "error"
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
        "fixed_interval": "<INTERVAL>m",
        "time_zone": "+08:00",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "aggregations": {
        "count": {
          "cardinality": {
            "field": "eid.keyword"
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

const errorList = `{
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
            "type": "error"
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
    "md5": {
      "terms": {
        "field": "md5.keyword",
        "size": 100
      },
      "aggregations": {
        "msg": {
          "terms": {
            "field": "error.msg.keyword"
          }
        },
        "error": {
          "terms": {
            "field": "error.error.keyword"
          }
        },
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        },
        "effectUser": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "ts": {
          "terms": {
            "field": "ts"
          }
        }
      }
    }
  }
}`

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

func (e elasticQuery) ErrorCountTrend(appId string, from, to, interval int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(errorCountTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "doc_count").Num
		eUser := gjson.Get(value.Raw, "uv.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()

		item := &response.CountItem{
			Count: int64(count),
			User:  int64(eUser),
			Key:   ts,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.CountListRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ErrorList(appId string, from, to int64) (*response.ErrorListRes, error) {
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
