package config

import "github.com/spf13/viper"

type SecretConfig struct {
	Secret string
}

func GetSecret() SecretConfig {
	return SecretConfig{
		Secret: viper.GetString("app.manage.secret"),
	}
}

type LogConfig struct {
	File string
}

func GetLog() LogConfig {
	return LogConfig{
		File: viper.GetString("app.manage.log.file"),
	}
}

type GormConfig struct {
	Driver string
	DSN    string
}

func GetGorm() GormConfig {
	return GormConfig{
		Driver: viper.GetString("app.manage.gorm.driver"),
		DSN:    viper.GetString("app.manage.gorm.dsn"),
	}
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func GetRedis() RedisConfig {
	return RedisConfig{
		Addr:     viper.GetString("app.manage.redis.addr"),
		Password: viper.GetString("app.manage.redis.password"),
		DB:       viper.GetInt("app.manage.redis.db"),
	}
}

type OssConfig struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	Secret    string
}

func GetOss() OssConfig {
	return OssConfig{
		Endpoint:  viper.GetString("app.manage.aliyun.endpoint"),
		Bucket:    viper.GetString("app.manage.aliyun.bucket"),
		AccessKey: viper.GetString("app.manage.aliyun.accessKey"),
		Secret:    viper.GetString("app.manage.aliyun.secret"),
	}
}

type MailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func GetMail() MailConfig {
	return MailConfig{
		Host:     viper.GetString("app.manage.mail.host"),
		Port:     viper.GetString("app.manage.mail.port"),
		Username: viper.GetString("app.manage.mail.username"),
		Password: viper.GetString("app.manage.mail.password"),
	}
}

type DingDingRobot struct {
	AccessToken string
	Secret      string
}

func GetRobot() DingDingRobot {
	return DingDingRobot{
		AccessToken: viper.GetString("app.manage.dingding.accessToken"),
		Secret:      viper.GetString("app.manage.dingding.secret"),
	}
}
