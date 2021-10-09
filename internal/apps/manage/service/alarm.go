package service

import (
	"context"
	"dora/internal/apps/manage/dao"
	"dora/internal/apps/manage/entity"
	"dora/internal/config"
	mailRes "dora/internal/datasource/mail"
	"dora/internal/datasource/redis"
	"dora/internal/logstore"
	"dora/pkg/utils"
	"dora/pkg/utils/dingTalk"
	"dora/pkg/utils/logx"
	"fmt"
	"strings"
	"time"
)

const (
	AlarmApiError  = "AlarmApiError"
	AlarmApiTimout = "AlarmApiTimout"
	AlarmJsError   = "AlarmJsError"
	AlarmResError  = "AlarmResError"
	AlarmPv        = "AlarmPv"
	AlarmUv        = "AlarmUv"
)

const (
	TimeUnitHour   = "hour"
	TimeUnitMinute = "minute"
)

const (
	ContactTypeEmail    = "email"
	ContactTypeDingTalk = "dingTalk"
	ContactTypeUser     = "user"
)

func ScanAllAlarms() {
	alarmDao := dao.NewAlarmDao()

	// todo 过滤 status
	list, err := alarmDao.List()
	if err != nil {
		logx.Errorf("scan alarms query list get err: %v", err)
		return
	}

	for _, alarm := range list {
		isAchieved, currentValue := checkAchieveCond(alarm)
		if isAchieved {
			sendAlarmMsg(alarm, currentValue)
		}
	}
}

// todo 其它规则实现
func checkAchieveCond(alarm *entity.Alarm) (bool, float64) {
	switch alarm.RuleType {
	case AlarmApiError:
		return false, 0

	case AlarmApiTimout:
		return false, 0

	case AlarmJsError:
		return checkAlarmJsError(alarm.AppId, alarm.Time, alarm.TimeUnit, alarm.Operator, alarm.Quota)

	case AlarmResError:
		return false, 0

	case AlarmPv:
		return false, 0

	case AlarmUv:
		return false, 0

	default:
		logx.Errorf("unknown alarm rule type %s", alarm.RuleType)
	}
	return false, 0
}

func checkAlarmJsError(appId string, timeVal int, timeUnit string, operator string, quota int) (bool, float64) {
	var unit time.Duration
	if timeUnit == TimeUnitHour {
		unit = time.Hour
	}
	if timeUnit == TimeUnitMinute {
		unit = time.Minute
	}
	period := time.Duration(timeVal) * unit

	f, t := utils.GetFormToRecently(period)
	count, err := logstore.GetClient().QueryMethods().ErrorCount(appId, f, t)
	if err != nil {
		logx.Errorf("corn alarm checkAchieveCond ErrorCount happen err: %v", err)
		return false, 0
	}

	// todo  使用 operator
	if count.Count >= quota {
		return true, float64(count.Count)
	}
	return false, 0
}

func sendAlarmMsg(alarm *entity.Alarm, nowValues float64) {
	contactDao := dao.NewAlarmContactDao()
	list, err := contactDao.List(alarm.ID, 1)
	if err != nil {
		logx.Errorf("corn alarm sendAlarmMsg contactDao.List err: %s", err)
		return
	}

	productDao := dao.NewProjectDao()
	cache := redis.Instance()

	for _, contact := range list {
		// 获取项目详情
		ProjectInfo, err := productDao.GetByAppId(alarm.AppId)
		if err != nil {
			logx.Errorf("corn alarm GetByAppId err: %s", err)
			continue
		}

		// 组装告警详情
		content := strings.Replace(alarm.Content, "{@num}", fmt.Sprintf("%v", nowValues), 1)
		message := fmt.Sprintf("%s 项目告警: %s 请及时修复！", ProjectInfo.Name, content)

		switch contact.Type {
		case ContactTypeUser:
			// todo 通知用户
		case ContactTypeEmail:
			if contact.Email == "" {
				break
			}
			// 静默标记检查
			key := fmt.Sprintf("alarm_email_silence_%v", alarm.ID)
			val := cache.Get(context.Background(), key).Val()
			if val != "" {
				break
			}

			// email 告警
			sendEmail(contact.Email, message)
			saveAlarmLogs(alarm, contact, message)

			// 设置静默标记
			cache.Set(context.Background(), key, message, time.Hour*3)

		case ContactTypeDingTalk:
			if contact.DingTalkAT == "" || contact.DingTalkSecret == "" {
				logx.Warnf("钉钉告警 access_token 或 secret 未配置，alarm id：%v", alarm.ID)
				break
			}

			// 静默标记检查
			key := fmt.Sprintf("alarm_dingtalk_silence_%v", alarm.ID)
			val := cache.Get(context.Background(), key).Val()
			if val != "" {
				break
			}

			// ding ding 告警
			sendDingTalk(contact.DingTalkAT, contact.DingTalkSecret, message)
			saveAlarmLogs(alarm, contact, message)

			// 设置静默标记
			cache.Set(context.Background(), key, message, time.Minute*5)
		default:
		}
	}
}

func sendEmail(toEmail string, message string) {
	conf := config.GetMail()

	to := toEmail
	fr := fmt.Sprintf("Dora Robot <%s>", conf.Username)
	sub := message
	con := message

	m := mailRes.BuilderEmail(to, fr, sub, con)
	err := mailRes.GetPool().Send(m, 10*time.Second)
	if err != nil {
		logx.Errorf("corn alarm dingTalk.SendDingDing err: %s", err)
		return
	}
	logx.Infof("success send alarm email : %s", message)
}

func sendDingTalk(accessToken, secret, message string) {
	dingMsg := dingTalk.NewDingTalkMsg(message)
	err := dingTalk.SendDingDing(dingMsg, accessToken, secret)
	if err != nil {
		logx.Errorf("corn alarm dingTalk.SendDingDing err: %s", err)
		return
	}
	logx.Infof("success send alarm dingtalk message : %s", message)
}

func saveAlarmLogs(alarm *entity.Alarm, contact *entity.AlarmContact, message string) {
	logDao := dao.NewAlarmLogDao()
	_, err := logDao.Create(&entity.AlarmLog{
		AlarmId:        alarm.ID,
		AlarmContactId: contact.ID,
		Content:        message,
	})
	if err != nil {
		logx.Errorf("corn alarm saveAlarmLogs err: %s", err)
		return
	}
}
