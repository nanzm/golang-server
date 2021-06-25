package boot

import (
	"dora/app/manage/datasource/gorm"
	"dora/app/manage/datasource/mail"
	"dora/app/manage/datasource/redis"
	"dora/app/manage/schedule"
	"dora/config"
	"dora/modules/initialize"
	"dora/modules/logstore/datasource/slslog"
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
	// 2：创建issues
	schedule.Cron()
}

func TearDown() {
	slslog.ClientTearDown()
	redis.StopClient()
	gorm.TearDown()
	mail.StopMailPool()
}
