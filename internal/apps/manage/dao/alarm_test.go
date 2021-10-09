package dao

import (
	"dora/internal/config"
	"dora/pkg/utils"
	"dora/pkg/utils/logx"
	"testing"
)

func init() {
	config.MustLoad("/Users/neil/Desktop/dora-platform/dora-server/config.local.yml")
	logx.Init("./dora.log")
}

func TestAlarm_List(t *testing.T) {
	dao := NewAlarmDao()
	list, err := dao.List()
	if err != nil {
		panic(err)
	}
	utils.PrettyPrint(list)
}
