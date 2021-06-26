package config

import (
	"dora/pkg/utils/fs"
	"errors"
	"github.com/spf13/viper"
	"log"
)

func init() {
	configPath := "/Users/neil/Desktop/dora-platform/dora-server/config.yml"
	//configPath := "config.yml"

	// file exist
	exists, err := fs.FileExists(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// load env file for local test
	if !exists {
		log.Fatal(errors.New("config path is wrong"))
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// load all Environment variables
	viper.AutomaticEnv()
}
