package core

import (
	"dora/modules/api/manage/config"
	"dora/modules/api/manage/schedule"
	"dora/modules/datasource/gorm"
	"dora/modules/datasource/mail"
	"dora/modules/datasource/redis"
	"dora/modules/datasource/slslog"
	"dora/modules/initialize"
	"dora/pkg/utils/logx"
)

func Setup() {
	// log
	conf := config.GetLog()
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
