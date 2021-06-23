package config

import (
	"dora/pkg/utils/fs"
	"dora/pkg/utils/logx"
	"github.com/spf13/viper"
)


func init() {
	//configPath := "/Users/neil/Desktop/dora-platform/dora-server/config.yml"
	configPath := "config.yml"

	// file exist
	exists, err := fs.FileExists(configPath)
	if err != nil {
		logx.Fatal(err)
	}

	// load env file for local test
	if exists {
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			logx.Fatal(err)
		}
	}

	// load all Environment variables
	viper.AutomaticEnv()
}
