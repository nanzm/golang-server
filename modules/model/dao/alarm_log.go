package dao

import (
	gorm2 "dora/modules/datasource/gorm"
	"dora/modules/model/dto"
	"dora/modules/model/entity"
	"gorm.io/gorm"
)

type AlarmLog struct {
	db *gorm.DB
}

func NewAlarmLogDao() *AlarmLog {
	return &AlarmLog{
		db: gorm2.GormInstance(),
	}
}

func (a *AlarmLog) Create(data *entity.AlarmLog) (*entity.AlarmLog, error) {
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
		Model(entity.AlarmLog{}).
		Delete(&entity.AlarmLog{ID: alarmId}).Error
	return err
}
