package config

import "github.com/spf13/viper"

type LogStore struct {
	Enable string
}

func GetLogStore() LogStore {
	return LogStore{
		Enable: viper.GetString("logStore.enable"),
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
		Endpoint:  viper.GetString("logStore.slsLog.endpoint"),
		AccessKey: viper.GetString("logStore.slsLog.accessKey"),
		Secret:    viper.GetString("logStore.slsLog.secret"),
		Project:   viper.GetString("logStore.slsLog.project"),
		LogStore:  viper.GetString("logStore.slsLog.logStore"),
		Topic:     viper.GetString("logStore.slsLog.topic"),
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
		Addresses: viper.GetStringSlice("logStore.elastic.addresses"),
		Username:  viper.GetString("logStore.elastic.username"),
		Password:  viper.GetString("logStore.elastic.password"),
		Index:     viper.GetString("logStore.elastic.index"),
	}
}
