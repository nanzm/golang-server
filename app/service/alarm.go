package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"dora/app/constant"
	"dora/app/dao"
	"dora/app/dto"
	"dora/app/logstore"
	"dora/pkg/logger"
	"dora/pkg/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Content struct {
	Content string `json:"content"`
}

type Payload struct {
	MsgType string  `json:"msgtype"`
	Text    Content `json:"text"`
}

func NewDingTalkMsg(content string) *Payload {
	return &Payload{
		MsgType: "text",
		Text: Content{
			Content: content,
		},
	}
}

func SendDingDing(payload *Payload, secret, accessToken string) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	signed, err := urlSign(timestamp, secret)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v&sign=%v&timestamp=%v", accessToken, signed, timestamp), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func urlSign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func alarmDingDing() {

}

func alarmEmail() {

}

func isAchieved(left int, Operator string, right int) bool {
	switch Operator {
	case constant.OperatorGt:
		return left > right
	case constant.OperatorLt:
		return left < right
	case constant.OperatorEq:
		return left == right
	case constant.OperatorGte:
		return left >= right
	case constant.OperatorLte:
		return left <= right
	default:
		logger.Error("未知操作符")
		return false
	}

}

func CornCheckAllProjectAlarm() {
	projectDao := dao.NewAlarmProjectDao()
	projectList, err := projectDao.List()
	if err != nil {
		logger.Error(err)
		return
	}

	for _, project := range projectList {
		if len(project.AlarmRules) > 0 && len(project.AlarmTargets) > 0 {
			matchRules(project)
		}
	}
}

// 检查是否满足规则
func matchRules(project dto.AlarmProject) {
	for _, rule := range project.AlarmRules {

		f, t := utils.GetFormToRecently(time.Minute * time.Duration(rule.Period))

		switch rule.Measure {
		case constant.MeasureApi:
			res, err := logstore.GetClient().QueryMethods().ApiErrorCount(project.ProjectInfo.AppId, f, t)
			if err != nil {
				logger.Error(err)
				return
			}
			if isAchieved(res.Count, rule.Operator, rule.Value) {
				msg := fmt.Sprintf("项目：%v api 错误数%v分钟内数量大于%v，共发生%v次，影响%v位用户",
					project.ProjectInfo.Name, rule.Period, rule.Value, res.Count, res.EffectUser)

				logger.Warnf("告警：%v", msg)
				triggerAlarm(msg, project)
			}

		case constant.MeasureError:
			res, err := logstore.GetClient().QueryMethods().ErrorCount(project.ProjectInfo.AppId, f, t)
			if err != nil {
				logger.Error(err)
				return
			}
			if isAchieved(res.Count, rule.Operator, rule.Value) {
				msg := fmt.Sprintf("项目：%v error 错误数%v分钟内数量大于%v，共发生%v次，影响%v位用户",
					project.ProjectInfo.Name, rule.Period, rule.Value, res.Count, res.EffectUser)

				logger.Warnf("告警：%v", msg)
				triggerAlarm(msg, project)
			}

		case constant.MeasureRes:
			res, err := logstore.GetClient().QueryMethods().ResLoadFailTotal(project.ProjectInfo.AppId, f, t)
			if err != nil {
				logger.Error(err)
				return
			}
			if isAchieved(res.Count, rule.Operator, rule.Value) {
				msg := fmt.Sprintf("项目：%v res 资源加载失败数%v分钟内数量大于%v，共发生%v次，影响%v位用户",
					project.ProjectInfo.Name, rule.Period, rule.Value, res.Count, res.EffectUser)

				logger.Warnf("告警：%v", msg)
				triggerAlarm(msg, project)
			}
		default:
			logger.Errorf("未实现的告警指标， 请修改")
		}
	}
}

func triggerAlarm(msg string, project dto.AlarmProject) {
}
