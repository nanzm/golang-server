package dao

import (
	"dora/internal/apps/manage/entity"
	dataGorm "dora/internal/datasource/gorm"
	"gorm.io/gorm"
)

type AlarmContactDao struct {
	db *gorm.DB
}

func NewAlarmContactDao() *AlarmContactDao {
	return &AlarmContactDao{
		db: dataGorm.Instance(),
	}
}

func (a *AlarmContactDao) Create(data *entity.AlarmContact) (*entity.AlarmContact, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AlarmContactDao) List(alarmId uint, status int) ([]*entity.AlarmContact, error) {
	list := make([]*entity.AlarmContact, 0)
	err := a.db.Model(&entity.AlarmContact{}).Where("alarm_id = ? AND status = ?", alarmId, status).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *AlarmContactDao) Update() {

}

//
//func (a *AlarmContactDao) Delete(alarmId uint) error {
//	err := a.db.
//		Model(model.AlarmProject{}).
//		Delete(&model.AlarmProject{ID: alarmId}).Error
//	return err
//}
