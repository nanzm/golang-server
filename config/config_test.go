package config

import (
	"dora/pkg/utils"
	"testing"
)

func init() {
	ParseConf("../config.yml")
}

func TestInitConfig(t *testing.T) {
	getConf := GetConf()
	utils.PrettyPrint(getConf)
}
