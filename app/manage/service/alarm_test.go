package service

import (
	"dora/pkg/utils/dingTalk"
	"fmt"
	"strings"
	"testing"
)

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
}

func TestFormat(t *testing.T) {
	var nowValues float64
	nowValues = 121
	content := strings.Replace("p0【脚本错误数量】 在 100 分钟内， 大于 10 次；当前为：{@num}", "{@num}",
		fmt.Sprintf("%v", nowValues), 1)

	fmt.Printf("%v \n", content)
}

func TestDingTalk(t *testing.T) {
	msg := dingTalk.NewDingTalkMsg("p0【脚本错误数量】 在 100 分钟内， 大于 10 次；当前为：{@num}")
	err := dingTalk.SendDingDing(msg, "cf50ae3eb9435c6172d4a2549c16f306f2a2e5026d2151414fc7318ca4c5aa4a", "SEC0d4930fc6b4145f63d83b29e26915a9c211bed43c58b9b37123a033f62673825")
	if err != nil {
		panic(err)
	}
}
