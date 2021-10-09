package boot

import (
	"dora/internal/apps/manage/schedule"
	"dora/internal/config"
	"dora/internal/datasource/gorm"
	"dora/internal/datasource/mail"
	"dora/internal/datasource/redis"
	"dora/internal/datasource/slslog"
	"dora/internal/initialize"
	"dora/pkg/utils/logx"
)

func Setup() {
	// log
	conf := config.GetManageLog()
	logx.Init(conf.File)

	// mail
	mail.GetPool()

	// redis
	redis.Instance()

	// database
	gorm.Instance()

	// 启动初始化
	initialize.Run()

	// 启动定时任务
	// 1：告警监控
	schedule.Cron()
}

func TearDown() {
	schedule.Stop()
	slslog.ClientTearDown()
	redis.StopClient()
	gorm.TearDown()
	mail.StopMailPool()
}
