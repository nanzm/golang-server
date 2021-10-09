package handler

import (
	"dora/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func HandleEvent(c *gin.Context, eventRaw []byte) (newEvents []byte, err error) {
	// 转 map
	var eventMaps []map[string]interface{}
	err = json.Unmarshal(eventRaw, &eventMaps)
	if err != nil {
		return nil, err
	}

	// 校验字段
	// 添加 ip
	// 添加 md5
	var ip = c.ClientIP()
	for _, eventMap := range eventMaps {
		err := verificationField(eventMap)
		if err != nil {
			return nil, err
		}
		eventMap["cip"] = ip

		md5 := genEventAggregateId(eventMap)
		if md5 != "" {
			eventMap["md5"] = md5
		}
	}

	// 转 byte
	result, err := json.Marshal(eventMaps)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func verificationField(event map[string]interface{}) error {
	if val, ok := event["type"]; !ok || val == "" {
		return errors.New("missing key \"type\"")
	}
	if val, ok := event["appId"]; !ok || val == "" {
		return errors.New("missing key \"appId\"")
	}
	return nil
}

// todo 更细致的聚合
func genEventAggregateId(event map[string]interface{}) string {
	val, ok := event["type"]
	if !ok || val == "" {
		return ""
	}

	data, ok := event["error"]
	if !ok || data == nil || data == "" {
		return ""
	}

	dataStr, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return utils.Md5(dataStr)
}
