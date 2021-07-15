package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

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
      "cardinality": {
        "field": "eid.keyword"
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
    "trend": {
      "date_histogram": {
        "field": "ts",
        "fixed_interval": "30m",
        "time_zone": "+08:00",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "aggregations": {
        "pv": {
          "cardinality": {
            "field": "eid.keyword"
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

const pageUrlPVUvList = `{
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
          "cardinality": {
            "field": "eid.keyword"
          }
        },
        "uv": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
		"bu": {
		  "cardinality": {
			"field": "buid.keyword"
		  }
		}
      }
    }
  }
}`

const projectLogCount = `{
  "size": 0,
  "query": {
    "bool": {
      "filter": [
        {
          "match": {
            "appId": "<APPID>"
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
    },
    "bu": {
      "cardinality": {
        "field": "buid.keyword"
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
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        }
      }
    }
  }
}`

const projectVersionList = `{
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
    "version": {
      "terms": {
        "field": "version.keyword",
        "size": 50,
        "order": {
          "count": "desc"
        }
      },
      "aggregations": {
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        }
      }
    }
  }
}`

const userScreenList = `{
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
    "screen": {
      "terms": {
        "field": "screen.keyword",
        "size": 50,
        "order": {
          "count": "desc"
        }
      },
      "aggregations": {
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        }
      }
    }
  }
}`

const logTypeList = `{
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
    "logType": {
      "terms": {
        "field": "type.keyword",
        "size": 50,
        "order": {
          "count": "desc"
        }
      },
      "aggregations": {
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        }
      }
    }
  }
}`

const projectEnvList = `{
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
    "env": {
      "terms": {
        "field": "env.keyword",
        "size": 50,
        "order": {
          "count": "desc"
        }
      },
      "aggregations": {
        "user": {
          "cardinality": {
            "field": "uid.keyword"
          }
        },
        "count": {
          "cardinality": {
            "field": "eid.keyword"
          }
        }
      }
    }
  }
}`

func (e elasticQuery) PvUvTotal(appId string, from, to int64) (*response.PvUvTotalRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(pvUvTotal, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	pv := gjson.Get(string(res), "aggregations.pv.value").Num
	uv := gjson.Get(string(res), "aggregations.uv.value").Num

	result := &response.PvUvTotalRes{
		Pv: int64(pv),
		Uv: int64(uv),
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
		pv := gjson.Get(value.Raw, "pv.value").Num
		uv := gjson.Get(value.Raw, "uv.value").Num
		ts := gjson.Get(value.Raw, "key_as_string").String()
		item := &response.PvUvTrendItemRes{
			Pv: int64(pv),
			Uv: int64(uv),
			Ts: ts,
		}
		logs = append(logs, item)
		return true // keep iterating
	})

	result := &response.PvUvTrendRes{
		Total: len(logs),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) PagesUrlVisitList(appId string, from, to int64) (*response.PageUrlVisitListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(pageUrlPVUvList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.PageUrlVisitItem, 0)
	buckets := gjson.Get(string(res), "aggregations.url.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		pv := gjson.Get(value.Raw, "pv.value").Num
		uv := gjson.Get(value.Raw, "uv.value").Num
		bu := gjson.Get(value.Raw, "bu.value").Num
		url := gjson.Get(value.Raw, "key").String()

		item := &response.PageUrlVisitItem{
			Pv:  int64(pv),
			Uv:  int64(uv),
			Bu:  int64(bu),
			Url: url,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.PageUrlVisitListRes{
		Total: int64(len(logs)),
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ProjectLogCount(appId string) (*response.ProjectLogCountRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(projectLogCount, appId, 0, 0))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	count := gjson.Get(string(res), "aggregations.count.value").Num
	user := gjson.Get(string(res), "aggregations.user.value").Num
	bu := gjson.Get(string(res), "aggregations.bu.value").Num

	result := &response.ProjectLogCountRes{
		Count: int64(count),
		User:  int64(user),
		Bu:    int64(bu),
	}
	return result, nil
}

func (e elasticQuery) ProjectLogTypeList(appId string, from, to int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(logTypeList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)
	buckets := gjson.Get(string(res), "aggregations.logType.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		version := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num

		item := &response.CountItem{
			Key:   version,
			Count: int64(count),
			User:  int64(user),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.CountListRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ProjectSdkVersionList(appId string, from, to int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(sdkVersionList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)
	buckets := gjson.Get(string(res), "aggregations.sdk.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		version := gjson.Get(value.Raw, "key").String()
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num

		item := &response.CountItem{
			Key:   version,
			Count: int64(count),
			User:  int64(user),
		}
		logs = append(logs, item)
		return true
	})

	result := &response.CountListRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ProjectEnvList(appId string, from, to int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(projectEnvList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)
	buckets := gjson.Get(string(res), "aggregations.env.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num
		env := gjson.Get(value.Raw, "key").String()

		item := &response.CountItem{
			Count: int64(count),
			User:  int64(user),
			Key:   env,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.CountListRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ProjectVersionList(appId string, from, to int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(projectVersionList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)
	buckets := gjson.Get(string(res), "aggregations.version.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num
		version := gjson.Get(value.Raw, "key").String()

		item := &response.CountItem{
			Count: int64(count),
			User:  int64(user),
			Key:   version,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.CountListRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}

func (e elasticQuery) ProjectUserScreenList(appId string, from, to int64) (*response.CountListRes, error) {
	res, err := baseSearch(e.config.Index, buildQueryTpl(userScreenList, appId, from, to))
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	logs := make([]*response.CountItem, 0)
	buckets := gjson.Get(string(res), "aggregations.screen.buckets")
	buckets.ForEach(func(key, value gjson.Result) bool {
		count := gjson.Get(value.Raw, "count.value").Num
		user := gjson.Get(value.Raw, "user.value").Num
		screen := gjson.Get(value.Raw, "key").String()

		item := &response.CountItem{
			Count: int64(count),
			User:  int64(user),
			Key:   screen,
		}
		logs = append(logs, item)
		return true
	})

	result := &response.CountListRes{
		Total: 0,
		List:  logs,
	}
	return result, nil
}
