package dao

import (
	"dora/app/constant"
	"dora/app/model"
	"testing"
)

func TestAlarmRuleDao_Create(t *testing.T) {
	dao := NewAlarmRuleDao()
	_, err := dao.Create(&model.AlarmRule{
		Name:       "错误太多",
		Period:     20,
		Measure:    "error",
		Aggregates: constant.AggregatesSum,
		Operator:   constant.OperatorEq,
		Value:      20,
	})

	if err != nil {
		panic(err)
	}
}
