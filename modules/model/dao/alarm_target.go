package dao

import (
	gorm2 "dora/modules/datasource/gorm"
	"gorm.io/gorm"
)

type AlarmTargetDao struct {
	db *gorm.DB
}

func NewAlarmTargetDao() *AlarmTargetDao {
	return &AlarmTargetDao{
		db: gorm2.GormInstance(),
	}
}

//
//func (a *AlarmTargetDao) Create(data *model.AlarmTarget) (*model.AlarmTarget, error) {
//	err := a.db.Create(data).Error
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func (a *AlarmTargetDao) Update() {
//
//}
//
//func (a *AlarmTargetDao) Delete(alarmId uint) error {
//	err := a.db.
//		Model(model.AlarmProject{}).
//		Delete(&model.AlarmProject{ID: alarmId}).Error
//	return err
//}
