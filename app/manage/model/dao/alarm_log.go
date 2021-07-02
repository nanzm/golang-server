package dao

import (
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"
	"gorm.io/gorm"
)

type AlarmLog struct {
	db *gorm.DB
}

func NewAlarmLogDao() *AlarmLog {
	return &AlarmLog{
		db: dataGorm.Instance(),
	}
}

func (a *AlarmLog) Create(data *entity.AlarmLog) (*entity.AlarmLog, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AlarmLog) List() (result []entity.AlarmLog, e error) {
	list := make([]entity.AlarmLog, 0)
	err := a.db.Debug().Model(entity.AlarmLog{}).
		Preload("AlarmProject").Find(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *AlarmLog) Delete(alarmId uint) error {
	err := a.db.
		Model(entity.AlarmLog{}).
		Delete(&entity.AlarmLog{ID: alarmId}).Error
	return err
}
