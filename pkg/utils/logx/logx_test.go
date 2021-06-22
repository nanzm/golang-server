package logx

import (
	"errors"
	"testing"
)


func TestStdout(t *testing.T) {
	Println("米好", "米好", "米好", "米好")
	Printf("%v info", "123123")
	Debugf("%v debug", "123123")
	Infof("%v info", "123123")
	Warnf("%v warn", "123123")

	Error(errors.New("123"))
	Errorf("%v error", "123123")

	//Panicf("%v panic", "123123")
}

func TestInfo(t *testing.T) {
	//Log.Info("你好")
	//Sugar.Info("你好")
	//Sugar.Info("你好", "你好")
	//Infof("%v 1", "1")
	//Println("123123123")
	//Printf("%v 1", "1231231")
	Warnf("%v 1", "1231231")
	Warn("hello", " world")
}

func TestMap(t *testing.T) {
	d := make(map[string]string)
	d["1"] = "hello"
	d["2"] = "hello"
	d["3"] = "hello"
	d["4"] = "hello"
	d["5"] = "hello"
	//sugar.Info("你好", d)
}
