// https://github.com/dustin/go-humanize
package utils

import (
	"fmt"
	"math"
	"sort"
	"time"
)

const (
	Day      = 24 * time.Hour
	Week     = 7 * Day
	Month    = 30 * Day
	Year     = 12 * Month
	LongTime = 37 * Year
)

type relTimeMagnitude struct {
	D      time.Duration
	Format string
	DivBy  time.Duration
}

var enMagnitudes = []relTimeMagnitude{
	{time.Second, "now", time.Second},
	{2 * time.Second, "1 second %s", 1},
	{time.Minute, "%d seconds %s", time.Second},
	{2 * time.Minute, "1 minute %s", 1},
	{time.Hour, "%d minutes %s", time.Minute},
	{2 * time.Hour, "1 hour %s", 1},
	{Day, "%d hours %s", time.Hour},
	{2 * Day, "1 day %s", 1},
	{Week, "%d days %s", Day},
	{2 * Week, "1 week %s", 1},
	{Month, "%d weeks %s", Week},
	{2 * Month, "1 month %s", 1},
	{Year, "%d months %s", Month},
	{18 * Month, "1 year %s", 1},
	{2 * Year, "2 years %s", 1},
	{LongTime, "%d years %s", Year},
	{math.MaxInt64, "a long while %s", 1},
}

var zhMagnitudes = []relTimeMagnitude{
	{time.Second, "现在", time.Second},
	{2 * time.Second, "1秒%s", 1},
	{time.Minute, "%d秒%s", time.Second},
	{2 * time.Minute, "1分钟%s", 1},
	{time.Hour, "%d分钟%s", time.Minute},
	{2 * time.Hour, "1小时%s", 1},
	{Day, "%d小时%s", time.Hour},
	{2 * Day, "1天%s", 1},
	{Week, "%d天%s", Day},
	{2 * Week, "1星期%s", 1},
	{Month, "%d星期%s", Week},
	{2 * Month, "1月%s", 1},
	{Year, "%d月%s", Month},
	{18 * Month, "1年%s", 1},
	{2 * Year, "2年%s", 1},
	{LongTime, "%d年%s", Year},
	{math.MaxInt64, "很久%s", 1},
}

func TimeFromNow(then time.Time) string {
	return relTime(then, time.Now(), "前", "后", zhMagnitudes)
}

func TimeFromNowEn(then time.Time) string {
	return relTime(then, time.Now(), "ago", "from now", enMagnitudes)
}

func relTime(a, b time.Time, prefix, suffix string, magnitudes []relTimeMagnitude) string {
	return customRelTime(a, b, prefix, suffix, magnitudes)
}

func customRelTime(a, b time.Time, prefix, suffix string, magnitudes []relTimeMagnitude) string {
	lbl := prefix
	diff := b.Sub(a)

	if a.After(b) {
		lbl = suffix
		diff = a.Sub(b)
	}

	n := sort.Search(len(magnitudes), func(i int) bool {
		return magnitudes[i].D > diff
	})

	if n >= len(magnitudes) {
		n = len(magnitudes) - 1
	}
	mag := magnitudes[n]
	var args []interface{}
	escaped := false
	for _, ch := range mag.Format {
		if escaped {
			switch ch {
			case 's':
				args = append(args, lbl)
			case 'd':
				args = append(args, diff/mag.DivBy)
			}
			escaped = false
		} else {
			escaped = ch == '%'
		}
	}
	return fmt.Sprintf(mag.Format, args...)
}
