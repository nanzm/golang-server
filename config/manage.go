package config

import "github.com/spf13/viper"


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

func GetGorm() GormConfig {
	return GormConfig{
		Driver: viper.GetString("app.manage.gorm.driver"),
		DSN:    viper.GetString("app.manage.gorm.dsn"),
	}
}

func GetRedis() RedisConfig {
	return RedisConfig{
		Addr:     viper.GetString("app.manage.redis.addr"),
		Password: viper.GetString("app.manage.redis.password"),
		DB:       viper.GetInt("app.manage.redis.db"),
	}
}

func GetOss() OssConfig {
	return OssConfig{
		Endpoint:  viper.GetString("app.manage.aliyun.endpoint"),
		Bucket:    viper.GetString("app.manage.aliyun.bucket"),
		AccessKey: viper.GetString("app.manage.aliyun.accessKey"),
		Secret:    viper.GetString("app.manage.aliyun.secret"),
	}
}

func GetMail() MailConfig {
	return MailConfig{
		Host:     viper.GetString("app.manage.mail.host"),
		Port:     viper.GetString("app.manage.mail.port"),
		Username: viper.GetString("app.manage.mail.username"),
		Password: viper.GetString("app.manage.mail.password"),
	}
}

func GetRobot() DingDingRobot {
	return DingDingRobot{
		AccessToken: viper.GetString("app.manage.dingding.accessToken"),
		Secret:      viper.GetString("app.manage.dingding.secret"),
	}
}
