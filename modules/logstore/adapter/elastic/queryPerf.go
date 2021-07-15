package elasticComponent

import (
	"dora/modules/logstore/response"
	"dora/pkg/utils/logx"
	"github.com/tidwall/gjson"
)

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
            "key": "<300",
            "to": 300
          },
          {
            "key": "300",
            "from": 300,
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
            "key": "<300",
            "to": 300
          },
          {
            "key": "300",
            "from": 300,
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
            "key": "<30",
            "to": 30
          },
          {
            "key": "30",
            "from": 30,
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
