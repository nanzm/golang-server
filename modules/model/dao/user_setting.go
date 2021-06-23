package dao

import (
	gorm2 "dora/modules/datasource/gorm"
	"dora/modules/model/dto"
	"dora/modules/model/entity"
	"gorm.io/gorm"
)

type UserSettingDao struct {
	db *gorm.DB
}

func NewUserSettingDao() *UserSettingDao {
	return &UserSettingDao{
		db: gorm2.GormInstance(),
	}
}

func (d *UserSettingDao) Create(userSetting *entity.UserSetting) (result *entity.UserSetting, error error) {
	err := d.db.Model(&entity.UserSetting{}).Create(userSetting).Error
	if err != nil {
		return nil, err
	}
	return userSetting, nil
}

func (d *UserSettingDao) Delete(userSettingId uint) error {
	err := d.db.Delete(&entity.UserSetting{ID: userSettingId}).Error
	return err
}

func (d *UserSettingDao) UpdateOrCreate(uid uint, projectId uint) error {
	q := entity.UserSetting{}
	err := d.db.Where("user_id = ?", uid).Find(&q).Error
	if err != nil {
		return err
	}

	// 没有 创建
	if q.UserId == 0 {
		c := entity.UserSetting{
			UserId:    uid,
			ProjectId: projectId,
		}
		err = d.db.Create(&c).Error
		if err != nil {
			return err
		}
		return nil
	}

	// 有 更新
	u := entity.UserSetting{ProjectId: projectId}
	err = d.db.Where("user_id = ?", uid).Select("project_id").Updates(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserSettingDao) Get(userId uint) (userSettingVO *dto.UserSettingVo, error error) {
	var p dto.UserSettingVo
	error = d.db.Where("user_id = ? ", userId).Preload("Project").Find(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *UserSettingDao) List(cur, size int) (
	result []entity.UserSetting, current, pageSize int, total int64, error error) {

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
	list := make([]entity.UserSetting, 0)

	err := d.db.Model(entity.UserSetting{}).Preload("Users").
		Count(&t).Limit(s).Offset((n - 1) * s).Order("id desc").Find(&list).Error

	if err != nil {
		return list, n, s, t, err
	}
	return list, n, s, t, nil
}
