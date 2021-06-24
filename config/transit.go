package config

import "github.com/spf13/viper"

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

func GetNsq() NsqConfig {
	return NsqConfig{
		Address: viper.GetString("app.transit.nsq.address"),
		Topic:   viper.GetString("app.transit.nsq.topic"),
		Channel: viper.GetString("app.transit.nsq.channel"),
	}
}
