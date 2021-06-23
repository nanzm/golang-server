package elastic

import (
	"dora/pkg/utils"
	"strings"
	"testing"
)

func TestGetElasticClient(t *testing.T) {
	es := GetElasticClient()
	info, err := es.Info()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info)
}

func TestPut(t *testing.T) {
	es := GetElasticClient()

	cp := `{
  "mappings": {
    "properties": {
      "category": {
        "type": "keyword"
      },
      "type": {
        "type": "keyword"
      },
      "ip": {
        "type": "keyword"
      },
      "breadcrumbs": {
        "type": "object",
        "enabled": false
      },
      "data": {
        "type": "object",
        "enabled": false
      },
      "error_response": {
        "type": "object",
        "enabled": false
      },
      "d_version": {
        "type": "keyword"
      },
      "d_url": {
        "type": "keyword"
      },
      "d_tit": {
        "type": "keyword"
      },
      "d_uuid": {
        "type": "keyword"
      },
      "d_ua_platform": {
        "type": "keyword"
      },
      "d_env": {
        "type": "keyword"
      },
      "d_s_wh": {
        "type": "keyword"
      },
      "d_ua_engine": {
        "type": "keyword"
      },
      "d_ua_os": {
        "type": "keyword"
      },
      "d_sdk_p": {
        "type": "keyword"
      },
      "d_sdk_v": {
        "type": "keyword"
      },
      "d_ua_browser": {
        "type": "keyword"
      },
      "d_eid": {
        "type": "keyword"
      },
      "d_ts": {
        "type": "date",
        "format": "epoch_second"
      },
      "d_appId": {
        "type": "keyword"
      },
      "d_ua": {
        "type": "keyword"
      },
      "d_send_mode": {
        "type": "keyword"
      }
    }
  }
}`
	mapping, err2 := es.Indices.Create("dora_test", es.Indices.Create.WithBody(strings.NewReader(cp)))

	if err2 != nil {
		t.Fatal(err2)
	}

	utils.PrettyPrint(mapping)
}
