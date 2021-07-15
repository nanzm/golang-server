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

const getLogsByMd5 = `{
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

const searchErrorLogs = `{
  "size": 100,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "error.error": "<SEARCH_ERROR>"
          }
        }
      ],
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
		Count:      int64(count),
		EffectUser: int64(effectUser),
		Logs:       logs,
	}
	return result, nil
}

func (e elasticQuery) SearchErrorLog(appId string, from, to int64, searchStr string) (*response.LogList, error) {
	r := strings.NewReplacer(
		core.TplAppId, appId,
		core.TplFrom, strconv.Itoa(int(from)),
		core.TplTo, strconv.Itoa(int(to)),
		core.TplSearchError, searchStr,
	)
	tpl := r.Replace(searchErrorLogs)

	res, err := baseSearch(e.config.Index, tpl)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	l := gjson.Get(string(res), "hits.hits").String()

	var logs []map[string]interface{}
	err = jsoniter.Unmarshal([]byte(l), &logs)
	if err != nil {
		return nil, err
	}

	result := &response.LogList{
		Total: len(logs),
		Logs:  logs,
	}
	return result, nil
}
