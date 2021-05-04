package dao

import (
	"dora/app/datasource"
	"dora/app/model"
	"gorm.io/gorm"
)

type AlarmRuleDao struct {
	db *gorm.DB
}

func NewAlarmRuleDao() *AlarmRuleDao {
	return &AlarmRuleDao{
		db: datasource.GormInstance(),
	}
}

func (a *AlarmRuleDao) Create(data *model.AlarmStrategy) (*model.AlarmStrategy, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AlarmRuleDao) Update() {

}
//
//func (a *AlarmRuleDao) Delete(alarmId uint) error {
//	err := a.db.
//		Model(model.AlarmProject{}).
//		Delete(&model.AlarmProject{ID: alarmId}).Error
//	return err
//}
