package dao

import (
	"dora/app/manage/model/dto"
	"dora/app/manage/model/entity"
	dataGorm "dora/modules/datasource/gorm"
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
		db: dataGorm.Instance(),
	}
}

func (d *UserDao) Create(user *entity.User) (result *entity.User, error error) {
	err := d.db.Model(&entity.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) Delete(userId uint) error {
	err := d.db.Delete(&entity.User{ID: userId}).Error
	return err
}

func (d *UserDao) Update(user *entity.User) error {
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

func (d *UserDao) UserProjects(userId uint) (result *entity.User, error error) {
	user := entity.User{}
	err := d.db.Model(&entity.User{}).Where("id = ?", userId).Preload("Projects").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
