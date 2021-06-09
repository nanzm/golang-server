package dao

import (
	"dora/internal/datasource"
	"dora/internal/dto"
	"gorm.io/gorm"
)

type AlarmProject struct {
	db *gorm.DB
}

func NewAlarmProjectDao() *AlarmProject {
	return &AlarmProject{
		db: datasource.GormInstance(),
	}
}
//
//func (a *AlarmProject) Create(data *model.AlarmProject) (*model.AlarmProject, error) {
//	err := a.db.Create(data).Error
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func (a *AlarmProject) AppendRules(alarmProjectId uint, rules []model.AlarmRule) error {
//	err := a.db.Model(&model.AlarmProject{ID: alarmProjectId}).Association("AlarmRules").Append(rules)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (a *AlarmProject) AppendTargets(alarmProjectId uint, targets []model.AlarmTarget) error {
//	err := a.db.Model(&model.AlarmProject{ID: alarmProjectId}).Association("AlarmTargets").Append(targets)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (a *AlarmProject) List() (result []dto.AlarmProject, e error) {
	list := make([]dto.AlarmProject, 0)
	err := a.db.Debug().Model(dto.AlarmProject{}).Where("silence > 0").
		Preload("ProjectInfo").
		Preload("AlarmRules").
		Preload("AlarmTargets").Find(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}
//
//func (a *AlarmProject) Delete(alarmId uint) error {
//	err := a.db.
//		Model(model.AlarmProject{}).
//		Delete(&model.AlarmProject{ID: alarmId}).Error
//	return err
//}
