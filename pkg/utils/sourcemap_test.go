package utils

import (
	"dora/pkg/logger"
	"testing"
)

func TestGetStackSourceMap(t *testing.T) {
	var s = `TypeError: Cannot read property 'refresh' of null\n    at p.refreshImmediately (https://ddd.xxx.com.cn/static/js/1.0563b6bd.chunk.js:1:419542)\n    at p.flush (https://ddd.xxx.com.cn/static/js/1.0563b6bd.chunk.js:1:419695)\n    at V.<anonymous> (https://ddd.xxx.com.cn/static/js/1.0563b6bd.chunk.js:1:1685)\n    at d (https://ddd.xxx.com.cn/static/js/1.0563b6bd.chunk.js:1:234284)",
	"url": "https://ddd.xxx.com.cn/static/js/1.0563b6bd.chunk.js`

	sourceMap, err := GetStackSourceMap("../tmp/sourcemap", s)
	if err != nil {
		panic(err)
	}
	logger.Printf("%s \n", sourceMap)
}
