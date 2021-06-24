package utils

import (
	jsoniter "github.com/json-iterator/go"
)

func StringToMap(rawData []byte) (map[string]interface{}, error) {
	var event map[string]interface{}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(rawData, &event)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func StringToMapList(rawData []byte) ([]map[string]interface{}, error) {
	var event []map[string]interface{}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(rawData, &event)

	if err != nil {
		return nil, err
	}

	return event, nil
}
