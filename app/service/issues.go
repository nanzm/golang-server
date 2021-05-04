package service

import (
	"context"
	"dora/app/constant"

	"dora/app/dao"
	"dora/app/datasource"
	"dora/app/logstore"
	"dora/app/model"
	"dora/config"
	"dora/pkg/logger"
	"dora/pkg/utils"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"time"
)

type IssuesService interface {
	CornCreateCheck()
	CornUpdateCheck()
	CreateIssues(slsLog map[string]string)
	QueryLogsGetCount(f, t int64, md5 string) (total, uCount int)
	GetAllMd5() []string
}

type issues struct {
	conf *config.Conf
}

func NewIssuesService() IssuesService {
	return &issues{
		conf: config.GetConf(),
	}
}

// 定时检查 创建
func (i *issues) CornCreateCheck() {
	ctx := context.Background()
	keys, err := datasource.RedisInstance().SMembers(ctx, constant.Md5ListWaitCreate).Result()
	if err != nil {
		logger.Errorf("CornCreateCheck redis get err: v%", err)
		return
	}

	if len(keys) > 0 {
		f, t := utils.GetDayFromNowRange(30)
		// 依次查出
		for _, md5 := range keys {
			logs, e := logstore.GetClient().QueryMethods().GetLogByMd5(f, t, md5)
			if e != nil {
				logger.Errorf("CornCreateCheck err: v%", e)
				return
			}

			// 查到对应的 log 创建 issues
			if len(logs.Logs) > 0 {
				i.CreateIssues(logs.Logs[0])
			} else {
				logger.Warn("该 md5 未找到对应的 log 无法创建 issues: ", md5)
			}
		}
	}
}

// 定时检查 更新
func (i *issues) CornUpdateCheck() {
	ctx := context.Background()
	keys, err := datasource.RedisInstance().SMembers(ctx, constant.Md5ListWaitUpdate).Result()
	if err != nil {
		logger.Error(err)
		return
	}

	if len(keys) > 0 {
		for _, md5 := range keys {
			i.UpdateIssues(md5)
		}
	}
}

// 创建 issues
func (i *issues) CreateIssues(slsLog map[string]string) {
	logger.Info("创建 Issues :", slsLog["agg_md5"])

	// 没有
	isu := &model.Issue{
		AppId:      slsLog["_appId"],
		AppVersion: slsLog["_version"],
		Env:        slsLog["_env"],
		Md5:        slsLog["agg_md5"],
		Type:       slsLog["type"],
		Category:   slsLog["category"],
		Raw:        utils.SafeJsonMarshal(slsLog),
		Resolve:    false,
		Ignore:     false,
	}

	issueDao := dao.NewIssueDao()
	_, err := issueDao.Create(isu)
	if err != nil {
		logger.Errorf("错误 CreateIssues issueDao.Create : %s", err)
		return
	}
	logger.Info("创建成功 Issues :", slsLog["agg_md5"])

	// 添加 md5 到已有列表
	datasource.RedisSetAdd(constant.Md5ListHas, slsLog["agg_md5"])

	// 删除这个 md5
	ctx := context.Background()
	_, err = datasource.RedisInstance().SRem(ctx, constant.Md5ListWaitCreate, slsLog["agg_md5"]).Result()
	if err != nil {
		logger.Error(err)
		return
	}
}

// 更新 issues
func (i *issues) UpdateIssues(md5 string) {
	logger.Info("更新 Issues :", md5)

	issueDao := dao.NewIssueDao()
	err := issueDao.UpdateByMd5(md5, &model.Issue{UpdatedAt: time.Now()})
	if err != nil {
		logger.Errorf("UpdateIssues failed: %s", err)
		return
	}
	logger.Info("更新成功 Issues :", md5)

	// 清空
	ctx := context.Background()
	_, err = datasource.RedisInstance().Del(ctx, constant.Md5ListWaitUpdate).Result()
	if err != nil {
		logger.Error(err)
		return
	}
}

// 根据 md5 查询日志出现多少次、uv 等
func (i *issues) QueryLogsGetCount(f, t int64, md5 string) (total, uCount int) {
	s := logstore.GetClient()
	md5Log, err := s.QueryMethods().LogCountByMd5(f, t, md5)
	if err != nil || md5Log == nil {
		return 0, 0
	}
	return md5Log.Count, md5Log.EffectUser
}

func (i *issues) GetAllMd5() []string {
	var md5s []string
	err := datasource.GormInstance().Model(&model.Issue{}).Select("md5").Find(&md5s).Error
	if err != nil {
		return nil
	}
	return md5s
}

func slsLogToMap(slsLog *sls.Log) map[string]string {
	tmp := map[string]string{}
	for _, content := range slsLog.Contents {
		tmp[*content.Key] = *content.Value
	}
	return tmp
}
