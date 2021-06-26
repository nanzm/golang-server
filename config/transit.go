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
