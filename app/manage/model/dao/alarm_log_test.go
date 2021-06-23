package dao

import (
	"dora/app/manage/model/entity"
	"dora/pkg/utils"
	"testing"
)

func TestAlarmLog_Create(t *testing.T) {
	dao := NewAlarmLogDao()
	_, err := dao.Create(&entity.AlarmLog{
		AlarmProjectId: 1,
		Log:            "hello world",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlarmLog_List(t *testing.T) {
	dao := NewAlarmLogDao()
	list, err := dao.List()
	if err != nil {
		t.Fatal(err)
	}

	utils.PrettyPrint(list)
}