package initialize

import (
	"dora/config"
	"dora/pkg/utils/logx"
	"testing"
)

func init() {
	config.MustLoad("/Users/neil/Desktop/dora-platform/dora-server/config.yml")
	logx.Init("./dora.log")
}

func Test_createDocMapping(t *testing.T) {
	createDocMapping()
}
