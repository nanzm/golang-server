package dao

import (
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"
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

func (a *Artifact) List() (result []*entity.Artifact, e error) {
	list := make([]*entity.Artifact, 0)
	err := a.db.Model(&entity.Artifact{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *Artifact) Delete(artifactId uint) error {
	err := a.db.
		Model(entity.Artifact{}).
		Delete(&entity.Artifact{ID: artifactId}).Error
	return err
}
