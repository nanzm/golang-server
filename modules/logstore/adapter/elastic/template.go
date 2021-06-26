package elasticComponent

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
            "to": 500
          },
          {
            "from": 500,
            "to": 1000
          },
          {
            "from": 1000,
            "to": 1500
          },
          {
            "from": 1500,
            "to": 2000
          },
          {
            "from": 2000,
            "to": 2500
          },
          {
            "from": 2500,
            "to": 3500
          },
          {
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
            "to": 500
          },
          {
            "from": 500,
            "to": 1000
          },
          {
            "from": 1000,
            "to": 1500
          },
          {
            "from": 1500,
            "to": 2000
          },
          {
            "from": 2000,
            "to": 2500
          },
          {
            "from": 2500,
            "to": 3500
          },
          {
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
            "to": 150
          },
          {
            "from": 150,
            "to": 300
          },
          {
            "from": 300,
            "to": 450
          },
          {
            "from": 450,
            "to": 600
          },
          {
            "from": 600,
            "to": 800
          },
          {
            "from": 800,
            "to": 1000
          },
          {
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
            "to": 0.1
          },
          {
            "from": 0.1,
            "to": 0.15
          },
          {
            "from": 0.15,
            "to": 0.2
          },
          {
            "from": 0.2,
            "to": 0.25
          },
          {
            "from": 0.25,
            "to": 0.3
          },
          {
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
            "to": 50
          },
          {
            "from": 100,
            "to": 150
          },
          {
            "from": 150,
            "to": 200
          },
          {
            "from": 200,
            "to": 300
          },
          {
            "from": 300,
            "to": 400
          },
          {
            "from": 400,
            "to": 500
          },
          {
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
            "to": 1000
          },
          {
            "from": 1000,
            "to": 1500
          },
          {
            "from": 1500,
            "to": 2000
          },
          {
            "from": 2000,
            "to": 2500
          },
          {
            "from": 2500,
            "to": 3000
          },
          {
            "from": 3000,
            "to": 4000
          },
          {
            "from": 4000,
            "to": 5000
          },
          {
            "from": 5000,
            "to": 6000
          },
          {
            "from": 6000
          }
        ]
      }
    }
  }
}`
