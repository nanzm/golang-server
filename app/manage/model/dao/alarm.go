package dao

import (
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"
	"dora/pkg/utils"
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

func (a *Alarm) List() (result []*entity.Alarm, e error) {
	list := make([]*entity.Alarm, 0)
	err := a.db.Model(&entity.Alarm{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}


func (a *Alarm) ListWithQuery(current, pageSize int64, appId string) (result []*entity.Alarm, Count int64, e error) {
	list := make([]*entity.Alarm, 0)
	var total int64

	db := a.db.Debug().Model(&entity.Alarm{})
	if appId != "" {
		db = db.Where("app_id = ?", appId)
	}
	err := db.
		Count(&total).
		Scopes(utils.Paginate(current, pageSize)).
		Find(&list).
		Order("id desc").Error

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
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
