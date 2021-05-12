package dao

import (
	"dora/app/datasource"
	"dora/app/dto"
	"dora/app/model"
	"encoding/gob"

	"gorm.io/gorm"
)

func init() {
	gob.Register(dto.UserWithRole{})
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{
		db: datasource.GormInstance(),
	}
}

func (d *UserDao) Create(user *model.User) (result *model.User, error error) {
	err := d.db.Model(&model.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) Delete(userId uint) error {
	err := d.db.Delete(&model.User{ID: userId}).Error
	return err
}

func (d *UserDao) Update(user *model.User) error {
	err := d.db.Model(user).Updates(user).Error
	return err
}

func (d *UserDao) Get(userId uint) (user *dto.UserWithRole, error error) {
	var p dto.UserWithRole
	p.ID = userId

	error = d.db.Preload("Role").First(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *UserDao) GetByEmail(email string) (user *dto.UserWithRole, error error) {
	var p dto.UserWithRole
	error = d.db.Preload("Role").Where("email = ?", email).Find(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *UserDao) GetByName(username string) (user *dto.UserWithRole, error error) {
	var p dto.UserWithRole
	error = d.db.Preload("Role").Where("username = ?", username).Find(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *UserDao) List(cur, size int) (
	result []dto.UserWithRole, current, pageSize int, total int64, error error) {

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
	list := make([]dto.UserWithRole, 0)

	err := d.db.Model(dto.UserWithRole{}).Preload("Role").
		Count(&t).Limit(s).Offset((n - 1) * s).Order("id desc").Find(&list).Error

	if err != nil {
		return list, n, s, t, err
	}
	return list, n, s, t, nil
}

func (d *UserDao) UserProjects(userId uint) (result *model.User, error error) {
	user := model.User{
		ID: userId,
	}
	err := d.db.Preload("Projects").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
