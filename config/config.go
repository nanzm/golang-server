package config

import (
	"dora/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type Conf struct {
	Debug  bool
	Secret string

	Gorm    GormConfig
	Redis   RedisConfig
	Nsq     NsqConfig
	Oss     OssConfig
	SlsLog  SlsLog
	Elastic Elastic
	Mail    MailConfig
}

var conf Conf

func GetConf() *Conf {
	if conf.Secret == "" {
		logger.Panic("please run \"ParseConf\" func to load config.yml!")
	}
	return &conf
}

var onceParseConf sync.Once

//  e.g., ParseConf("./config.yml")
func ParseConf(path string) *Conf {
	onceParseConf.Do(func() {
		viper.SetConfigFile(path)

		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			logger.Errorf("err: %v", err)
		}
		logger.Infof("config file read success: %v", viper.ConfigFileUsed())

		if err := viper.Unmarshal(&conf); err != nil {
			logger.Panicf("unmarshal yaml config failed: %v", err)
		}
	})

	return &conf
}
