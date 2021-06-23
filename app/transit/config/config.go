package config

import "github.com/spf13/viper"

type SecretConfig struct {
	Secret string
}

func GetSecret() SecretConfig {
	return SecretConfig{
		Secret: viper.GetString("app.transit.secret"),
	}
}

type LogConfig struct {
	File string
}

func GetLog() LogConfig {
	return LogConfig{
		File: viper.GetString("app.transit.log.file"),
	}
}

type NsqConfig struct {
	Address string
	Topic   string
	Channel string
}

func GetNsq() NsqConfig {
	return NsqConfig{
		Address: viper.GetString("app.transit.nsq.address"),
		Topic:   viper.GetString("app.transit.nsq.topic"),
		Channel: viper.GetString("app.transit.nsq.channel"),
	}
}
