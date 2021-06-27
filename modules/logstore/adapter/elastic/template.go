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

const errorList=`{
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
          "value_count": {
            "field": "type.keyword"
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
          "value_count": {
            "field": "type.keyword"
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