package task

import (
	"dora/app/service"
	"dora/pkg/logger"
	"github.com/robfig/cron/v3"
)

func StartCron() {
	c := cron.New(cron.WithSeconds())
	registerAlarmCheck(c)
	registerIssuesCheck(c)

	c.Start()

	//t1 := time.NewTimer(time.Second * 10)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}

func registerAlarmCheck(c *cron.Cron) {
	_, err := c.AddFunc("0/1 * * * * ? ", func() {

	})

	if err != nil {
		logger.Printf("registerAlarmCheck cron %v:", err)
	}

}

func registerIssuesCheck(c *cron.Cron) {
	_, err := c.AddFunc("0/5 * * * * ? ", func() {
		issuesService := service.NewIssuesService()
		issuesService.CornCreateCheck()
		issuesService.CornUpdateCheck()
	})

	if err != nil {
		logger.Printf("registerIssuesCheck cron %v:", err)
	}
}
