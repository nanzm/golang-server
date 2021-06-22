package config

import "github.com/spf13/viper"

type AppConfig struct {
	Secret string
}

func GetApp() AppConfig {
	return AppConfig{
		Secret: viper.GetString("secret"),
	}
}

type GormConfig struct {
	Driver string
	DSN    string
}

func GetGorm() GormConfig {
	return GormConfig{
		Driver: viper.GetString("gorm.driver"),
		DSN:    viper.GetString("gorm.dsn"),
	}
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func GetRedis() RedisConfig {
	return RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
}

type NsqConfig struct {
	Address string
	Topic   string
	Channel string
}

func GetNsq() NsqConfig {
	return NsqConfig{
		Address: viper.GetString("nsq.address"),
		Topic:   viper.GetString("nsq.topic"),
		Channel: viper.GetString("nsq.channel"),
	}
}

type SlsLog struct {
	Endpoint  string
	AccessKey string
	Secret    string
	Project   string
	LogStore  string
	Topic     string
	Source    string
}

func GetSlsLog() SlsLog {
	return SlsLog{
		Endpoint:  viper.GetString("slsLog.endpoint"),
		AccessKey: viper.GetString("slsLog.accessKey"),
		Secret:    viper.GetString("slsLog.secret"),
		Project:   viper.GetString("slsLog.project"),
		LogStore:  viper.GetString("slsLog.logStore"),
		Topic:     viper.GetString("slsLog.topic"),
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
		Endpoint:  viper.GetString("aliyun.endpoint"),
		Bucket:    viper.GetString("aliyun.bucket"),
		AccessKey: viper.GetString("aliyun.accessKey"),
		Secret:    viper.GetString("aliyun.secret"),
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
		Host:     viper.GetString("mail.host"),
		Port:     viper.GetString("mail.port"),
		Username: viper.GetString("mail.username"),
		Password: viper.GetString("mail.password"),
	}
}

type Elastic struct {
	Addresses []string
	Username  string
	Password  string
	Index     string
}

func GetElastic() Elastic {
	return Elastic{
		Addresses: viper.GetStringSlice("elastic.addresses"),
		Username:  viper.GetString("elastic.username"),
		Password:  viper.GetString("elastic.password"),
		Index:     viper.GetString("elastic.index"),
	}
}

type DingDingRobot struct {
	AccessToken string
	Secret      string
}

func GetRobot() DingDingRobot {
	return DingDingRobot{
		AccessToken: viper.GetString("dingding.accessToken"),
		Secret:      viper.GetString("dingding.secret"),
	}
}
