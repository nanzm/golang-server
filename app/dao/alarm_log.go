package dao

import (
	"dora/app/datasource"
	"dora/app/dto"
	"dora/app/model"
	"gorm.io/gorm"
)

type AlarmLog struct {
	db *gorm.DB
}

func NewAlarmLogDao() *AlarmLog {
	return &AlarmLog{
		db: datasource.GormInstance(),
	}
}

func (a *AlarmLog) Create(data *model.AlarmLog) (*model.AlarmLog, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AlarmLog) List() (result []dto.AlarmLog, e error) {
	list := make([]dto.AlarmLog, 0)
	err := a.db.Debug().Model(dto.AlarmLog{}).
		Preload("AlarmProject").Find(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *AlarmLog) Delete(alarmId uint) error {
	err := a.db.
		Model(model.AlarmLog{}).
		Delete(&model.AlarmLog{ID: alarmId}).Error
	return err
}
