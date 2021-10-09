package schedule

import (
	"dora/internal/apps/manage/service"
	"dora/pkg/utils/logx"
	"github.com/robfig/cron/v3"
)

var c *cron.Cron

func Cron() {
	c = cron.New(cron.WithSeconds())
	registerAlarmCheck(c)

	c.Start()
}

func registerAlarmCheck(c *cron.Cron) {
	_, err := c.AddFunc("0/5 * * * * ? ", func() {
		service.ScanAllAlarms()
	})

	if err != nil {
		logx.Fatalf("registerAlarmCheck cron %v:", err)
	}
}

func Stop() {
	if c != nil {
		c.Stop()
		logx.Infof("schedule corn has stop")
	}
}
