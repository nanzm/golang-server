package service

import (
	"dora/config"
	"testing"
)

func init() {
	config.ParseConf("../../config.yml")
}

func TestSendDingDingMsg(t *testing.T) {
	//conf := config.GetConf()
	//
	//logx.Println(conf.DingDing)
	//if len(conf.DingDing) == 0 {
	//	return
	//}
	//
	//secret := conf.DingDing[0].Secret
	//accessToken := conf.DingDing[0].AccessToken
	//
	//data := NewDingTalkMsg("测试 [鲜花]")
	//err := SendDingDing(data, secret, accessToken)
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestCornCheckAllProjectAlarm(t *testing.T) {
	CornCheckAllProjectAlarm()
}
