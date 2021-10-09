package config

import "github.com/spf13/viper"

const (
	// 上传文件路劲
	UploadDir = "storage/upload"

	// 项目制品 文件路径
	BackupDir = "storage/backup"

	// sourcemap
	SourcemapDir = "storage/sourcemap"
)

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
