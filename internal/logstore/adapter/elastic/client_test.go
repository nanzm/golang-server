package elasticComponent

import (
	"dora/internal/datasource/elastic"
	"dora/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func Test_elkPutErrorData(t *testing.T) {
	raw := `{"_eid":"86402155-e35f-4247-9a50-cf3b3f6acaef","_appId":"fca5deec-a9db-4dac-a4db-b0f4610d16a5","_version":"2021-03-29 15:21:41 +0800 @c613a558638b6a92e46b28525f49a0089917d790 @hide-180-gc613a55","_env":"production","_ts":1617024680,"_sdk_v":"1.0.1","_sdk_p":"browser","_tit":"杭州中小学数智校园","_url":"http://114.55.166.248/demo","_ua":"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36","_ua_browser":"Chrome 89.0.4389.90","_ua_engine":"Blink","_ua_os":"macOS 11.2.1","_ua_platform":"desktop Apple","_uuid":"275c5fcd-d787-4aa1-9014-d66b28a0dd5a","_s_wh":"1920×1080","_send_mode":"sendBeacon","agg":"{\"error\":\"Cannot read property 'a' of undefined\",\"stack\":\"TypeError: Cannot read property 'a' of undefined\\n    at onClick (http://114.55.166.248/static/js/44.6004e6e1.chunk.js:1:932)\\n    at Object.u (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:613741)\\n    at h (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:613884)\\n    at http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:614030\\n    at g (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:614116)\\n    at at (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:629719)\\n    at it (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:629529)\\n    at ut (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:629885)\\n    at ht (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:631093)\\n    at L (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:728245)\"}","breadcrumbs":[{"type":"click","timestamp":1617024650,"message":"click node","data":{"tagName":"BUTTON","id":"","className":"","name":"","outerHTML":"<button>按钮</button>","nodeType":1,"selector":"html > body > div#root > div.dashboard > div#dashboard.dashboard-content > div.sc-bdfBwQ cIKpxU > button:nth-child(1)"}},{"type":"console","timestamp":1617024650,"message":"console.log","data":["----------------d1221emo----------------"]},{"type":"click","timestamp":1617024658,"message":"click node","data":{"tagName":"BUTTON","id":"","className":"","name":"","outerHTML":"<button>按钮</button>","nodeType":1,"selector":"html > body > div#root > div.dashboard > div#dashboard.dashboard-content > div.sc-bdfBwQ cIKpxU > button:nth-child(1)"}},{"type":"console","timestamp":1617024658,"message":"console.log","data":["----------------d1221emo----------------"]},{"type":"click","timestamp":1617024679,"message":"click node","data":{"tagName":"DIV","id":"dashboard","className":"dashboard-content","outerHTML":"<div class=\"dashboard-content\" id=\"dashboard\"><div class=\"sc-bdfBwQ cIKpxU\"><h1>1</h1><button>按钮</button><div><code>1</code></div></div></div>","nodeType":1,"selector":"html > body > div#root > div.dashboard > div#dashboard.dashboard-content:nth-child(1)"}},{"type":"click","timestamp":1617024679,"message":"click node","data":{"tagName":"DIV","id":"","className":"sc-bdfBwQ cIKpxU","outerHTML":"<div class=\"sc-bdfBwQ cIKpxU\"><h1>1</h1><button>按钮</button><div><code>1</code></div></div>","nodeType":1,"selector":"html > body > div#root > div.dashboard > div#dashboard.dashboard-content > div.sc-bdfBwQ cIKpxU:nth-child(0)"}},{"type":"click","timestamp":1617024680,"message":"click node","data":{"tagName":"BUTTON","id":"","className":"","name":"","outerHTML":"<button>按钮</button>","nodeType":1,"selector":"html > body > div#root > div.dashboard > div#dashboard.dashboard-content > div.sc-bdfBwQ cIKpxU > button:nth-child(1)"}},{"type":"console","timestamp":1617024680,"message":"console.log","data":["----------------d1221emo----------------"]}],"category":"error","type":"onerror","msg":"Uncaught TypeError: Cannot read property 'a' of undefined","url":"http://114.55.166.248/static/js/12.2fc124a0.chunk.js","line":2,"column":629921,"error":"Cannot read property 'a' of undefined","stack":"TypeError: Cannot read property 'a' of undefined\n    at onClick (http://114.55.166.248/static/js/44.6004e6e1.chunk.js:1:932)\n    at Object.u (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:613741)\n    at h (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:613884)\n    at http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:614030\n    at g (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:614116)\n    at at (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:629719)\n    at it (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:629529)\n    at ut (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:629885)\n    at ht (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:631093)\n    at L (http://114.55.166.248/static/js/12.2fc124a0.chunk.js:2:728245)"}`
	toMap, err2 := utils.StringToMap([]byte(raw))

	if err2 != nil {
		t.Fatal(err2)
	}

	err := NewElkLogStore().PutData(toMap)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_elkSearch(t *testing.T) {
	es := elastic.GetClient()

	res, err := es.Search(
		es.Search.WithIndex("dora"),
		es.Search.WithBody(strings.NewReader(`{
    "size": 0,
    "aggs": {
        "category": {
            "terms": {
                "field": "category"
            }
        },
        "sdk": {
            "terms": {
                "field": "d_sdk_v"
            }
        },
        "avg_fid": {
            "avg": {
                "field": "fid"
            }
        }
    }
}`)),
	)

	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	r := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	utils.PrettyPrint(r)
}

func Test_addPrefix(t *testing.T) {
	//testData := map[string]interface{}{
	//	"a":  "12",
	//	"_b": "12",
	//}

	//result := addPrefix(testData, "d")
	//assert.Equal(t, result, map[string]interface{}{
	//	"a":   "12",
	//	"d_b": "122",
	//}, "they should be equal")
}

func Test_removePrefix(t *testing.T) {
	//testData := map[string]interface{}{
	//	"a":   "12",
	//	"d_b": "12",
	//}
	//
	//result := removePrefix(testData, "d")
	//assert.Equal(t, result, map[string]interface{}{
	//	"a":  "12",
	//	"_b": "12",
	//}, "they should be equal")
}

func Test_buildQueryTpl(t *testing.T) {
	fr, to := utils.GetFormToRecently(48 * time.Hour)
	appId := "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
	tpl := buildQueryTpl(searchErrorLogs, appId, fr, to)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(tpl)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func Test_buildQueryTrendTpl(t *testing.T) {
	var interval int64
	fr, to := utils.GetFormToRecently(60 * time.Hour)
	appId := "fca5deec-a9db-4dac-a4db-b0f4610d16a5"
	interval = 30


	tpl := buildQueryTrendTpl(resLoadFailTrend, appId, fr, to, interval)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(tpl)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}
