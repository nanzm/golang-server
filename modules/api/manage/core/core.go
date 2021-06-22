package core

import (
	"dora/modules/api/manage/schedule"
	"dora/modules/datasource"
	"dora/modules/initialize"
)

func Setup()  {
	// mail
	datasource.GetMailPool()

	// redis
	datasource.RedisInstance()

	// database
	datasource.GormInstance()

	// 启动初始化
	initialize.Run()

	// 启动定时任务
	// 1：告警监控
	// 2：创建issues
	schedule.Cron()
}

func TearDown()  {
	datasource.StopSlsLog()
	datasource.StopRedisClient()
	datasource.StopDataBase()
	datasource.StopMailPool()
}