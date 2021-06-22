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