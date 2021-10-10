package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	// 上传文件路劲
	UploadDir = "storage/upload"

	// 项目制品 文件路径
	BackupDir = "storage/backup"

	// sourcemap
	SourcemapDir = "storage/sourcemap"
)

var Uptime = time.Now()

func GetSlsLog() SlsLog {
	return SlsLog{
		Endpoint:  viper.GetString("datasource.logStore.slsLog.endpoint"),
		AccessKey: viper.GetString("datasource.logStore.slsLog.accessKey"),
		Secret:    viper.GetString("datasource.logStore.slsLog.secret"),
		Project:   viper.GetString("datasource.logStore.slsLog.project"),
		LogStore:  viper.GetString("datasource.logStore.slsLog.logStore"),
		Topic:     viper.GetString("datasource.logStore.slsLog.topic"),
	}
}

func GetElastic() Elastic {
	return Elastic{
		Addresses: viper.GetStringSlice("datasource.logStore.elastic.addresses"),
		Username:  viper.GetString("datasource.logStore.elastic.username"),
		Password:  viper.GetString("datasource.logStore.elastic.password"),
		Index:     viper.GetString("datasource.logStore.elastic.index"),
	}
}

func GetGorm() GormConfig {
	return GormConfig{
		Driver: viper.GetString("DORA_GORM_DRIVER"),
		DSN:    viper.GetString("DORA_GORM_DSN"),
	}
}

func GetRedis() RedisConfig {
	return RedisConfig{
		Addr:     viper.GetString("datasource.redis.addr"),
		Password: viper.GetString("datasource.redis.password"),
		DB:       viper.GetInt("datasource.redis.db"),
	}
}

func GetOss() OssConfig {
	return OssConfig{
		Endpoint:  viper.GetString("datasource.aliyun.endpoint"),
		Bucket:    viper.GetString("datasource.aliyun.bucket"),
		AccessKey: viper.GetString("datasource.aliyun.accessKey"),
		Secret:    viper.GetString("datasource.aliyun.secret"),
	}
}

func GetMail() MailConfig {
	return MailConfig{
		Host:     viper.GetString("datasource.mail.host"),
		Port:     viper.GetString("datasource.mail.port"),
		Username: viper.GetString("datasource.mail.username"),
		Password: viper.GetString("datasource.mail.password"),
	}
}

func GetDingTalkRobot() DingDingRobot {
	return DingDingRobot{
		AccessToken: viper.GetString("datasource.dingding.accessToken"),
		Secret:      viper.GetString("datasource.dingding.secret"),
	}
}

func GetNsq() NsqConfig {
	return NsqConfig{
		Address: viper.GetString("datasource.nsq.address"),
		Topic:   viper.GetString("datasource.nsq.topic"),
		Channel: viper.GetString("datasource.nsq.channel"),
	}
}

func GetManageSecret() SecretConfig {
	return SecretConfig{
		Secret: viper.GetString("app.manage.secret"),
	}
}

func GetManageLog() LogConfig {
	return LogConfig{
		File: viper.GetString("app.manage.log.file"),
	}
}

func GetTransitSecret() SecretConfig {
	return SecretConfig{
		Secret: viper.GetString("app.transit.secret"),
	}
}

func GetTransitLog() LogConfig {
	return LogConfig{
		File: viper.GetString("app.transit.log.file"),
	}
}
