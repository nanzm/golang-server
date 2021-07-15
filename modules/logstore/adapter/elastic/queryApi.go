package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

// 接口错误总数
const apiErrorCount = `{
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
            "type": "api"
          }
        },
        {
          "match": {
            "subType": "xhr"
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

const apiErrorTrend = `{
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
            "type": "api"
          }
        },
        {
          "match": {
            "subType": "xhr"
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
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        }
      }
    }
  }
}`

// 接口响应错误列表
const apiErrorList = `{
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
            "type": "api"
          }
        },
        {
          "match": {
            "subType": "xhr"
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
        "field": "api.url.keyword",
        "size": 50,
        "order": {
          "count": "desc"
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
        },
        "method": {
          "terms": {
            "field": "api.method.keyword"
          }
        },
        "type": {
          "terms": {
            "field": "api.type.keyword"
          }
        }
      }
    }
  }
}`

// 接口请求量排名前50 的响应时间
const apiTopListDuration = `{
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
            "type": "performance"
          }
        },
        {
          "match": {
            "subType": "resource"
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
        "field": "performance.xmlhttprequest.name.keyword",
        "size": 50,
        "order": {
          "user": "desc"
        }
      },
      "aggregations": {
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        },
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "xhr_percent": {
          "percentiles": {
            "field": "performance.xmlhttprequest.duration"
          }
        }
      }
    }
  }
}`

// 单个 url
const apiDuration = `{
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
            "type": "performance"
          }
        },
        {
          "match": {
            "subType": "resource"
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
    "xhr": {
      "range": {
        "field": "performance.xmlhttprequest.duration",
        "ranges": [
          {
            "key": "<100",
            "to": 100
          },
          {
            "key": "100",
            "from": 100,
            "to": 200
          },
          {
            "key": "200",
            "from": 200,
            "to": 300
          },
          {
            "key": "300",
            "from": 300,
            "to": 400
          },
          {
            "key": "400",
            "from": 400,
            "to": 500
          },
          {
            "key": "500",
            "from": 500,
            "to": 600
          },
          {
            "key": "600",
            "from": 600,
            "to": 700
          },
          {
            "key": "700",
            "from": 700,
            "to": 800
          },
          {
            "key": "800",
            "from": 800,
            "to": 900
          },
          {
            "key": "900",
            "from": 900,
            "to": 1000
          },
          {
            "key": "1000",
            "from": 1000,
            "to": 2000
          },
          {
            "key": "2000",
            "from": 2000,
            "to": 3000
          },
          {
            "key": "3000",
            "from": 3000,
            "to": 6000
          },
          {
            "key": ">6000",
            "from": 6000
          }
        ]
      }
    },
    "xhr_percent": {
      "percentiles": {
        "field": "performance.xmlhttprequest.duration"
      }
    }
  }
}`

// 单个 url 趋势
const apiDurationTrend = `{
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
            "type": "performance"
          }
        },
        {
          "match": {
            "subType": "resource"
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
        "xhr_percent": {
          "percentiles": {
            "field": "performance.xmlhttprequest.duration"
          }
        }
      }
    }
  }
}`

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

func (e elasticQuery) ApiErrorTrend(appId string, from, to int64, interval int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(apiErrorTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()

		item := &response.CountItem{
			Count: int64(count),
			User:  int64(user),
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

func (e elasticQuery) ApiTopListDuration(appId string, from, to int64) (*response.ApiTopListDurationRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(apiTopListDuration, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ApiTopItem, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(_, value gjson.Result) bool {
		key := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num

		Percent1 := gjson.Get(value.Raw, "xhr_percent.values.1\\.0").Num
		Percent5 := gjson.Get(value.Raw, "xhr_percent.values.5\\.0").Num
		Percent25 := gjson.Get(value.Raw, "xhr_percent.values.25\\.0").Num
		Percent50 := gjson.Get(value.Raw, "xhr_percent.values.50\\.0").Num
		Percent75 := gjson.Get(value.Raw, "xhr_percent.values.75\\.0").Num
		Percent95 := gjson.Get(value.Raw, "xhr_percent.values.95\\.0").Num
		Percent99 := gjson.Get(value.Raw, "xhr_percent.values.99\\.0").Num

		item := &response.ApiTopItem{
			Key:   key,
			Count: int64(count),
			User:  int64(user),

			Percent1:  Percent1,
			Percent5:  Percent5,
			Percent25: Percent25,
			Percent50: Percent50,
			Percent75: Percent75,
			Percent95: Percent95,
			Percent99: Percent99,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ApiTopListDurationRes{
		Total: len(logs),
		List:  logs,
	}

	return result, err
}

func (e elasticQuery) ApiDuration(appId string, from, to int64) (*response.ApiDurationRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(apiDuration, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ApiDurationItemRes, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.xhr.buckets")
	buckets.ForEach(func(_, value gjson.Result) bool {
		key := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "doc_count").Num

		item := &response.ApiDurationItemRes{
			Key:   key,
			Count: int(count),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ApiDurationRes{
		List: logs,
		Percent: &response.ApiDurationPercent{
			Percent1:  gjson.Get(string(res), "aggregations.xhr_percent.values.1\\.0").Num,
			Percent5:  gjson.Get(string(res), "aggregations.xhr_percent.values.5\\.0").Num,
			Percent25: gjson.Get(string(res), "aggregations.xhr_percent.values.25\\.0").Num,
			Percent50: gjson.Get(string(res), "aggregations.xhr_percent.values.50\\.0").Num,
			Percent75: gjson.Get(string(res), "aggregations.xhr_percent.values.75\\.0").Num,
			Percent95: gjson.Get(string(res), "aggregations.xhr_percent.values.95\\.0").Num,
			Percent99: gjson.Get(string(res), "aggregations.xhr_percent.values.99\\.0").Num,
		},
	}

	return result, err
}

func (e elasticQuery) ApiDurationTrend(appId string, from, to int64, interval int64) (*response.ApiDurationTrendRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(apiDurationTrend, appId, from, to, interval))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ApiDurationTrendItemRes, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.trend.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		Percent1 := gjson.Get(value.Raw, "xhr_percent.values.1\\.0").Num
		Percent5 := gjson.Get(value.Raw, "xhr_percent.values.5\\.0").Num
		Percent25 := gjson.Get(value.Raw, "xhr_percent.values.25\\.0").Num
		Percent50 := gjson.Get(value.Raw, "xhr_percent.values.50\\.0").Num
		Percent75 := gjson.Get(value.Raw, "xhr_percent.values.75\\.0").Num
		Percent95 := gjson.Get(value.Raw, "xhr_percent.values.95\\.0").Num
		Percent99 := gjson.Get(value.Raw, "xhr_percent.values.99\\.0").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()

		item := &response.ApiDurationTrendItemRes{
			Percent1:  Percent1,
			Percent5:  Percent5,
			Percent25: Percent25,
			Percent50: Percent50,
			Percent75: Percent75,
			Percent95: Percent95,
			Percent99: Percent99,
			Ts:        ts,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.ApiDurationTrendRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}
