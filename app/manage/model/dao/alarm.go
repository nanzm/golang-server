package dao

import (
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"
	"gorm.io/gorm"
)

type Alarm struct {
	db *gorm.DB
}

func NewAlarmDao() *Alarm {
	return &Alarm{
		db: dataGorm.Instance(),
	}
}

func (a *Alarm) Create(data *entity.Alarm) (*entity.Alarm, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

//todo query status === 0
func (a *Alarm) List() (result []*entity.Alarm, e error) {
	list := make([]*entity.Alarm, 0)
	err := a.db.Model(&entity.Alarm{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

//func (a *Alarm) Status(alarmId uint) error {
//	err := a.db.
//		Model(model.Alarm{}).
//		Delete(&model.Alarm{ID: alarmId}).Error
//	return err
//}

//func (a *Alarm) Delete(alarmId uint) error {
//	err := a.db.
//		Model(model.Alarm{}).
//		Delete(&model.Alarm{ID: alarmId}).Error
//	return err
//}
