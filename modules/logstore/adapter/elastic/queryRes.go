package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

// 资源加载失败统计 img script 等等
const resLoadFailCount = `{
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
}`

const resLoadFailTrend = `{
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

const resLoadFailList = `{
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

// 资源加载排名前50 的响应时间
const resTopListDuration = `{
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
        "field": "performance.script.name.keyword",
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
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "script_percent": {
          "percentiles": {
            "field": "performance.script.duration"
          }
        }
      }
    }
  }
}`

// 单个资源 url
const resDuration = `{
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

// 单个资源 url 趋势
const resDurationTrend = `{
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

func (e elasticQuery) ResLoadFailTotal(appId string, from, to int64) (*response.ResLoadFailTotalRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(resLoadFailCount, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	c := gjson.Get(string(res), "aggregations.count.value").Num
	u := gjson.Get(string(res), "aggregations.user.value").Num

	result := &response.ResLoadFailTotalRes{
		Count: int(c),
		User:  int(u),
	}
	return result, nil
}

func (e elasticQuery) ResLoadFailTrend(appId string, from, to, interval int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTrendTpl(resLoadFailTrend, appId, from, to, interval))
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

func (e elasticQuery) ResLoadFailList(appId string, from, to int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(resLoadFailList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)

	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num
		url := gjson.Get(value.Raw, "key").String()

		item := &response.CountItem{
			Count: int64(count),
			User:  int64(user),
			Key:   url,
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

func (e elasticQuery) ResTopListDuration(appId string, from, to int64) (*response.ResTopListDurationRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(resTopListDuration, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.ResTopItem, 0)

	// 遍历
	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(_, value gjson.Result) bool {
		key := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num

		Percent1 := gjson.Get(value.Raw, "script_percent.values.1\\.0").Num
		Percent5 := gjson.Get(value.Raw, "script_percent.values.5\\.0").Num
		Percent25 := gjson.Get(value.Raw, "script_percent.values.25\\.0").Num
		Percent50 := gjson.Get(value.Raw, "script_percent.values.50\\.0").Num
		Percent75 := gjson.Get(value.Raw, "script_percent.values.75\\.0").Num
		Percent95 := gjson.Get(value.Raw, "script_percent.values.95\\.0").Num
		Percent99 := gjson.Get(value.Raw, "script_percent.values.99\\.0").Num

		item := &response.ResTopItem{
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

	result := &response.ResTopListDurationRes{
		Total: len(logs),
		List:  logs,
	}

	return result, err
}

func (e elasticQuery) ResDuration(appId string, from, to int64) (*response.ResDurationRes, error) {
	panic("implement me")
}

func (e elasticQuery) ResDurationTrend(appId string, from, to, interval int64) (*response.CountListRes, error) {
	panic("implement me")
}

