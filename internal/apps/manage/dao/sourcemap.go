package dao

import (
	"dora/internal/apps/manage/entity"
	dataGorm "dora/internal/datasource/gorm"
	"dora/pkg/utils"
	"gorm.io/gorm"
)

type Sourcemap struct {
	db *gorm.DB
}

func NewSourcemapDao() *Sourcemap {
	return &Sourcemap{
		db: dataGorm.Instance(),
	}
}

func (a *Sourcemap) Create(data *entity.Sourcemap) (*entity.Sourcemap, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Sourcemap) List(current, pageSize int64, appId string) (result []*entity.Sourcemap, Count int64, e error) {
	list := make([]*entity.Sourcemap, 0)
	var total int64

	db := a.db.Debug().Model(&entity.Sourcemap{})
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

func (a *Sourcemap) Delete(sourcemapId uint) error {
	err := a.db.
		Model(entity.Sourcemap{}).
		Delete(&entity.Sourcemap{ID: sourcemapId}).Error
	return err
}
