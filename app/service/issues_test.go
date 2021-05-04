package service

import (
	"dora/app/dao"
	"dora/pkg/logger"
	"dora/pkg/utils"
	"testing"
	"time"
)

func Test_issues_CreateIssues(t *testing.T) {
	service := NewIssuesService()
	service.CornCreateCheck()
	time.Sleep(time.Second * 10000)
	//q := "* and md5: "
}

func Test_queryCount(t *testing.T) {

	d := dao.NewIssueDao()
	list, _, _, total, err := d.
		ListQueryTimeRange("fca5deec-a9db-4dac-a4db-b0f4610d16a5", 1612195200, 1612249776, 1, 10)
	if err != nil {
		panic(err)
	}
	logger.Println(total)

	// 遍历查询 count
	issuesService := NewIssuesService()
	for _, issue := range list {
		count, uCount := issuesService.QueryLogsGetCount(1612195200, 1612249776, issue.Md5)
		issue.EventCount = count
		issue.UserCount = uCount
	}

	utils.PrettyPrint(list)
}

func Test_issues_GetAllMd5(t *testing.T) {
	service := NewIssuesService()
	md5 := service.GetAllMd5()
	logger.Printf("%v \n", len(md5))
	logger.Printf("%v \n", md5)
}
