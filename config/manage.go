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
