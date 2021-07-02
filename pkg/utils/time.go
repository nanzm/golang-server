package utils

import (
	"time"
)

const sysTimeFmt = "2006-01-02 15:04:05"
const sysTimeFmtShort = "2006-01-02"

var sysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")

// 获取时间范围 unix millis
func GetFromToRange(begin time.Time, howLong time.Duration) (from, to int64) {
	f := begin.Add(-howLong).UnixNano() / 1000000
	t := begin.UnixNano() / 1000000
	return f, t
}

// 获取最近 unix millis
func GetFormToRecently(howLong time.Duration) (from, to int64) {
	f := time.Now().Add(-howLong).UnixNano() / 1000000
	t := time.Now().UnixNano() / 1000000
	return f, t
}

// 获取多少天的范围 包括从 n 天前 0点开始 unix millis
func GetDayFromNowRange(n int) (from, to int64) {
	now := time.Now()
	f := now.Add(-(time.Hour * 24 * time.Duration(n)))
	fStart := time.Date(f.Year(), f.Month(), f.Day(), 0, 0, 0, 0, f.Location()).UnixNano() / 1000000
	t := now.UnixNano() / 1000000
	return fStart, t
}

func CurrentTime() string {
	format := time.Now().Format(sysTimeFmt)
	return format
}
