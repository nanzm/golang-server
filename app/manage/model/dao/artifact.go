package dao

import (
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"
	"dora/pkg/utils"
	"gorm.io/gorm"
)

type Artifact struct {
	db *gorm.DB
}

func NewArtifactDao() *Artifact {
	return &Artifact{
		db: dataGorm.Instance(),
	}
}

func (a *Artifact) Create(data *entity.Artifact) (*entity.Artifact, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Artifact) List(current, pageSize int64, appId, fileType string) (result []*entity.Artifact, Count int64, e error) {
	list := make([]*entity.Artifact, 0)
	var total int64

	db := a.db.Model(&entity.Artifact{})
	if appId != "" {
		db = db.Where("app_id = ?", appId)
	}
	if fileType != "" {
		db = db.Where("file_type = ?", fileType)
	}
	err := db.
		Scopes(utils.Paginate(current, pageSize)).
		Find(&list).
		Count(&total).
		Order("id desc").Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (a *Artifact) Delete(artifactId uint) error {
	err := a.db.
		Model(entity.Artifact{}).
		Delete(&entity.Artifact{ID: artifactId}).Error
	return err
}
