package dao

import (
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"

	"gorm.io/gorm"
)

type ProjectDao struct {
	db *gorm.DB
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		db: dataGorm.Instance(),
	}
}

func (d *ProjectDao) Create(project *entity.Project, uid uint) (result *entity.Project, error error) {
	// 创建
	err := d.db.Model(&entity.Project{}).Create(project).Error
	if err != nil {
		return nil, err
	}

	// 关联用户
	var createUser entity.User
	err = d.db.Model(&entity.User{}).Where("id = ?", uid).Find(&createUser).Error
	if err != nil {
		return nil, err
	}
	err = d.db.Model(&project).Association("Users").Append(&createUser)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (d *ProjectDao) Delete(projectId uint) error {
	err := d.db.Delete(&entity.Project{ID: projectId}).Error
	return err
}

// 忽略 非0值
func (d *ProjectDao) Update(project entity.Project) error {
	err := d.db.Model(&project).Updates(&project).Error
	return err
}

func (d *ProjectDao) Get(projectId uint) (project *entity.Project, error error) {
	var p entity.Project
	p.ID = projectId

	error = d.db.First(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *ProjectDao) GetByName(projectName string) (project *entity.Project, error error) {
	var p entity.Project
	error = d.db.Where("name=?", projectName).Find(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *ProjectDao) List(cur, size int) (
	result []entity.Project, current, pageSize int, total int64, error error) {

	// 默认1
	n := 1
	if cur > 0 {
		n = cur
	}

	// 默认10
	s := 10
	if size > 0 {
		s = size
	}

	var t int64
	list := make([]entity.Project, 0)

	err := d.db.Model(entity.Project{}).
		Count(&t).Limit(s).Offset((n - 1) * s).Order("id desc").Find(&list).Error

	if err != nil {
		return list, n, s, t, err
	}
	return list, n, s, t, nil
}

func (d *ProjectDao) ProjectUsers(projectId uint) (projects []*entity.Project, error error) {
	list := make([]*entity.Project, 0)
	error = d.db.Where("id=?", projectId).Preload("Users").Find(&list).Error
	if error != nil {
		return nil, error
	}
	return list, nil
}
