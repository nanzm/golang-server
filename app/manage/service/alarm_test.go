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
	//err := SendDingDing(data, accessToken, secret)
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
	err := dingTalk.SendDingDing(msg, "4dd273b36ba4df74618f71ed91baca18228ee26323cd9aa0b436b2835ac767d7", "SEC83c260bea45c5647c837024fec8b87424596469c140ae7976109c82bf3f8f3f3")
	if err != nil {
		panic(err)
	}
}
