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

func NewElasticQuery() core.Query {
	return &elasticQuery{
		config: config.GetElastic(),
	}
}


const pvUvTotal = `{
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
            "type": "visit"
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
    "uv": {
      "cardinality": {
        "field": "uid.keyword"
      }
    },
    "pv": {
      "value_count": {
        "field": "type.keyword"
      }
    }
  }
}`

const pvUvTotalTrend = `{
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
            "type": "visit"
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
    "pv": {
      "value_count": {
        "field": "type.keyword"
      }
    },
    "trend": {
      "date_histogram": {
        "field": "ts",
        "interval": "<INTERVAL>m",
        "time_zone": "+08:00",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "aggregations": {
        "uv": {
          "cardinality": {
            "field": "uid.keyword"
          }
        }
      }
    }
  }
}`

const urlPVUv = `{
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
            "type": "visit"
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
        "field": "href.keyword",
        "size": 50,
        "order": {
          "pv": "desc"
        }
      },
      "aggregations": {
        "pv": {
          "value_count": {
            "field": "type.keyword"
          }
        },
        "uv": {
          "cardinality": {
            "field": "uid.keyword"
          }
        }
      }
    }
  }
}`

const performanceBucket = `{
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
            "subType": "metric"
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
    "fp": {
      "range": {
        "field": "performance.fp",
        "ranges": [
          {
            "key": "<500",
            "to": 500
          },
          {
            "key": "500",
            "from": 500,
            "to": 1000
          },
          {
            "key": "1000",
            "from": 1000,
            "to": 1500
          },
          {
            "key": "1500",
            "from": 1500,
            "to": 2000
          },
          {
            "key": "2000",
            "from": 2000,
            "to": 2500
          },
          {
            "key": "2500",
            "from": 2500,
            "to": 3000
          },
          {
            "key": "3000",
            "from": 3000,
            "to": 3500
          },
          {
            "key": ">3500",
            "from": 3500
          }
        ]
      }
    },
    "fcp": {
      "range": {
        "field": "performance.fcp",
        "ranges": [
          {
            "key": "<500",
            "to": 500
          },
          {
            "key": "500",
            "from": 500,
            "to": 1000
          },
          {
            "key": "1000",
            "from": 1000,
            "to": 1500
          },
          {
            "key": "1500",
            "from": 1500,
            "to": 2000
          },
          {
            "key": "2000",
            "from": 2000,
            "to": 2500
          },
          {
            "key": "2500",
            "from": 2500,
            "to": 3000
          },
          {
            "key": "3000",
            "from": 3000,
            "to": 3500
          },
          {
            "key": ">3500",
            "from": 3500
          }
        ]
      }
    },
    "ttfb": {
      "range": {
        "field": "performance.ttfb",
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
            "key": ">1000",
            "from": 1000
          }
        ]
      }
    },
    "cls": {
      "range": {
        "field": "performance.cls",
        "ranges": [
          {
            "key": "<0.1",
            "to": 0.1
          },
          {
            "key": "0.1",
            "from": 0.1,
            "to": 0.15
          },
          {
            "key": "0.15",
            "from": 0.15,
            "to": 0.2
          },
          {
            "key": "0.2",
            "from": 0.2,
            "to": 0.25
          },
          {
            "key": "0.25",
            "from": 0.25,
            "to": 0.3
          },
          {
            "key": ">0.3",
            "from": 0.3
          }
        ]
      }
    },
    "fid": {
      "range": {
        "field": "performance.fid",
        "ranges": [
          {
            "key": "<50",
            "to": 50
          },
          {
            "key": "50",
            "from": 50,
            "to": 100
          },
          {
            "key": "100",
            "from": 100,
            "to": 150
          },
          {
            "key": "150",
            "from": 150,
            "to": 200
          },
          {
            "key": "200",
            "from": 200,
            "to": 250
          },
          {
            "key": "250",
            "from": 250,
            "to": 300
          },
          {
            "key": "300",
            "from": 300,
            "to": 350
          },
          {
            "key": "350",
            "from": 350,
            "to": 400
          },
          {
            "key": "400",
            "from": 400,
            "to": 450
          },
          {
            "key": "450",
            "from": 450,
            "to": 500
          },
          {
            "key": ">500",
            "from": 500
          }
        ]
      }
    },
    "lcp": {
      "range": {
        "field": "performance.lcp",
        "ranges": [
          {
            "key": "<500",
            "to": 500
          },
          {
            "key": "1000",
            "from": 1000,
            "to": 1500
          },
          {
            "key": "1500",
            "from": 1500,
            "to": 2000
          },
          {
            "key": "2000",
            "from": 2000,
            "to": 2500
          },
          {
            "key": "2500",
            "from": 2500,
            "to": 3000
          },
          {
            "key": "3000",
            "from": 3000,
            "to": 3500
          },
          {
            "key": "3500",
            "from": 3500,
            "to": 4000
          },
          {
            "key": "4000",
            "from": 4000,
            "to": 4500
          },
          {
            "key": "4500",
            "from": 4500,
            "to": 5000
          },
          {
            "key": "5000",
            "from": 5000,
            "to": 5500
          },
          {
            "key": "5500",
            "from": 5500,
            "to": 6000
          },
          {
            "key": ">6000",
            "from": 6000
          }
        ]
      }
    }
  }
}`

const sdkVersionList = `{
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
    "sdk": {
      "terms": {
        "field": "sdk.keyword",
        "size": 50,
        "order": {
          "count": "desc"
        }
      },
      "aggregations": {
        "count": {
          "value_count": {
            "field": "appId.keyword"
          }
        }
      }
    }
  }
}`

const getLogsByMd5=`{
  "size": 100,
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
            "md5": "<MD5>"
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
}`

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
