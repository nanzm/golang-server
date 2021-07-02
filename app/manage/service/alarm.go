package service

import (
	"dora/app/manage/model/dao"
	"dora/app/manage/model/entity"
	"dora/config"
	mailRes "dora/modules/datasource/mail"
	"dora/modules/logstore"
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
		} else {
			logx.Infof("alarm scan %v not achieved, current: %v", alarm.RuleType, currentValue)
		}
	}
}

// todo 其它规则是实现
func checkAchieveCond(alarm *entity.Alarm) (bool, float64) {
	switch alarm.RuleType {
	case AlarmApiError:
		return false, 0

	case AlarmApiTimout:
		return false, 0

	case AlarmJsError:
		var unit time.Duration
		if alarm.TimeUnit == TimeUnitHour {
			unit = time.Hour
		}
		if alarm.TimeUnit == TimeUnitMinute {
			unit = time.Minute
		}
		period := time.Duration(alarm.Time) * unit

		f, t := utils.GetFormToRecently(period)
		count, err := logstore.GetClient().QueryMethods().ErrorCount(alarm.AppId, f, t)
		if err != nil {
			logx.Errorf("corn alarm checkAchieveCond ErrorCount happen err: %v", err)
			return false, 0
		}
		// todo  使用 alarm.Operator
		if count.Count >= alarm.Quota {
			return true, float64(count.Count)
		}

		return false, float64(count.Count)

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

func sendAlarmMsg(alarm *entity.Alarm, nowValues float64) {
	contactDao := dao.NewAlarmContactDao()
	list, err := contactDao.List(alarm.ID, 1)
	if err != nil {
		logx.Errorf("corn alarm sendAlarmMsg contactDao.List err: %s", err)
		return
	}

	// todo 通知用户  按 p0 p1 分类通知
	for _, contact := range list {
		replacedContent := strings.Replace(alarm.Content, "{@num}", fmt.Sprintf("%v", nowValues), 1)

		if contact.Type == ContactTypeUser {
		}

		if contact.Type == ContactTypeEmail {
			if contact.Email != "" {
				conf := config.GetMail()

				to := contact.Email
				fr := fmt.Sprintf("Dora Robot <%s>", conf.Username)
				sub := replacedContent
				con := replacedContent

				m := mailRes.BuilderEmail(to, fr, sub, con)
				err := mailRes.GetPool().Send(m, 10*time.Second)
				if err != nil {
					logx.Errorf("corn alarm dingTalk.SendDingDing err: %s", err)
					return
				}
				saveAlarmLogs(alarm, contact, replacedContent)
			}
		}

		if contact.Type == ContactTypeDingTalk {
			fmt.Println("------------------------------------")
			if contact.DingTalkAT != "" && contact.DingTalkSecret != "" {
				msg := dingTalk.NewDingTalkMsg(replacedContent)
				err := dingTalk.SendDingDing(msg, contact.DingTalkAT, contact.DingTalkSecret)
				if err != nil {
					logx.Errorf("corn alarm dingTalk.SendDingDing err: %s", err)
					return
				}
				saveAlarmLogs(alarm, contact, replacedContent)
			}
		}
	}
}

func saveAlarmLogs(alarm *entity.Alarm, contact *entity.AlarmContact, content string) {
	logDao := dao.NewAlarmLogDao()
	_, err := logDao.Create(&entity.AlarmLog{
		AlarmId:        alarm.ID,
		AlarmContactId: contact.ID,
		Content:        content,
	})
	if err != nil {
		logx.Errorf("corn alarm saveAlarmLogs err: %s", err)
		return
	}
}
