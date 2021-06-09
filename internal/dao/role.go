package dao

import (
	"dora/internal/datasource"
	"dora/internal/model"

	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao() *RoleDao {
	return &RoleDao{
		db: datasource.GormInstance(),
	}
}

func (d *RoleDao) Create(role *model.Role) (result *model.Role, error error) {
	err := d.db.Model(&model.Role{}).Create(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (d *RoleDao) Delete(roleId uint) error {
	err := d.db.Delete(&model.Role{ID: roleId}).Error
	return err
}

func (d *RoleDao) Update(role model.Role) error {
	err := d.db.Model(&role).Updates(&role).Error
	return err
}

func (d *RoleDao) Get(roleId uint) (role *model.Role, error error) {
	var p model.Role
	p.ID = roleId

	error = d.db.Preload("Users").First(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *RoleDao) List(cur, size int) (
	result []model.Role, current, pageSize int, total int64, error error) {

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
	list := make([]model.Role, 0)

	err := d.db.Model(model.Role{}).Preload("Users").
		Count(&t).Limit(s).Offset((n - 1) * s).Order("id desc").Find(&list).Error

	if err != nil {
		return list, n, s, t, err
	}
	return list, n, s, t, nil
}
