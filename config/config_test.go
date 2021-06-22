package config

import (
	"dora/pkg/utils"
	"testing"
)

func TestGetElastic(t *testing.T) {
	f := GetElastic()
	utils.PrettyPrint(f)
}

func TestGetGorm(t *testing.T) {
	f := GetGorm()
	utils.PrettyPrint(f)
}

func TestGetMail(t *testing.T) {
	f := GetMail()
	utils.PrettyPrint(f)
}

func TestGetNsq(t *testing.T) {
	f := GetNsq()
	utils.PrettyPrint(f)
}

func TestGetOss(t *testing.T) {
	f := GetOss()
	utils.PrettyPrint(f)
}

func TestGetRedis(t *testing.T) {
	f := GetRedis()
	utils.PrettyPrint(f)
}

func TestGetRobot(t *testing.T) {
	f := GetRobot()
	utils.PrettyPrint(f)
}

func TestGetSlsLog(t *testing.T) {
	f := GetSlsLog()
	utils.PrettyPrint(f)
}