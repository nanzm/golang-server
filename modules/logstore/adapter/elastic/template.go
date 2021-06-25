package elasticComponent

const pvUvTotal = `{
  "size": 0,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "d_appId": "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
          }
        },
        {
          "match": {
            "category": "view"
          }
        },
        {
          "range": {
            "d_ts": {
              "gte": "<FORM>",
              "lte": "<TO>"
            }
          }
        }
      ]
    }
  },
  "aggs": {
    "uv": {
      "cardinality": {
        "field": "d_uuid"
      }
    },
    "pv": {
      "value_count": {
        "field": "category"
      }
    }
  }
}`

const pvUvTotalTrend = `{
  "size": 0,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "d_appId": "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
          }
        },
        {
          "match": {
            "category": "view"
          }
        },
        {
          "range": {
            "d_ts": {
              "gte": "<FORM>",
              "lte": "<TO>"
            }
          }
        }
      ]
    }
  },
  "aggs": {
    "pv": {
      "value_count": {
        "field": "category"
      }
    },
    "pvTrend": {
      "date_histogram": {
        "field": "d_ts",
        "fixed_interval": "<INTERVAL>m",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "aggs": {
        "uv": {
          "cardinality": {
            "field": "d_uuid"
          }
        }
      }
    }
  }
}`

const entryPage = `{
  "size": 0,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "d_appId": "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
          }
        },
        {
          "match": {
            "category": "view"
          }
        },
        {
          "range": {
            "d_ts": {
              "gte": "<FORM>",
              "lte": "<TO>"
            }
          }
        }
      ]
    }
  },
  "aggs": {
    "entryPage": {
      "terms": {
        "field": "d_url",
        "size": 50,
        "order": {
          "pv": "desc"
        }
      },
      "aggs": {
        "pv": {
          "value_count": {
            "field": "category"
          }
        },
        "uv": {
          "cardinality": {
            "field": "d_uuid"
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
      "must": [
        {
          "match": {
            "d_appId": "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
          }
        },
        {
          "match": {
            "category": "error"
          }
        },
        {
          "range": {
            "d_ts": {
              "gte": "<FORM>",
              "lte": "<TO>"
            }
          }
        }
      ]
    }
  },
  "aggs": {
    "uv": {
      "cardinality": {
        "field": "d_uuid"
      }
    },
    "pv": {
      "value_count": {
        "field": "category"
      }
    }
  }
}`

const performanceBucket = `{
  "size": 0,
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "appId": "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
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
              "gte": "<FORM>",
              "lte": "<TO>"
            }
          }
        }
      ]
    }
  },
  "aggs": {
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
