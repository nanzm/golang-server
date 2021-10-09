package config

import (
	"dora/pkg/utils/fs"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func MustLoad(configPath string) {
	// file exist
	exists, err := fs.FileExists(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// load env file for local test
	if !exists {
		log.Fatal(errors.New(fmt.Sprintf("config is not exists: %s", configPath)))
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// load all Environment variables
	viper.AutomaticEnv()
}
