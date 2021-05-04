package dao

import (
	"dora/app/datasource"
	"dora/app/model"
	"dora/pkg/logger"
	"gorm.io/gorm"
)

type OrganizationDao struct {
	db *gorm.DB
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		db: datasource.GormInstance(),
	}
}

func (d *OrganizationDao) Create(uid uint, organization *model.Organization) (result *model.Organization, error error) {
	err := d.db.Model(&model.Organization{}).Create(organization).Error
	if err != nil {
		return nil, err
	}
	// 关联用户
	err = d.AddUser(organization.ID, []uint{uid})
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (d *OrganizationDao) GetMembers(organizationId uint) (members []*model.User, error error) {
	var p model.Organization
	p.ID = organizationId

	userList := make([]*model.User, 0)
	err := d.db.Model(&p).Association("Users").Find(&userList)
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (d *OrganizationDao) AddUser(organizationId uint, userIds []uint) error {
	userList := make([]*model.User, 0)
	for _, Id := range userIds {
		userList = append(userList, &model.User{
			ID: Id,
		})
	}
	logger.Println("------------------------------------")
	logger.Printf("%v \n", userList)
	logger.Println("------------------------------------")

	err := d.db.Model(&model.Organization{ID: organizationId}).Association("Users").
		Append(userList)
	return err
}

func (d *OrganizationDao) RemoveUser(organizationId, userId uint) error {
	err := d.db.Model(&model.Organization{ID: organizationId}).Association("Users").
		Delete(&model.User{ID: userId})
	return err
}

func (d *OrganizationDao) Delete(organizationId uint) error {
	err := d.db.Delete(&model.Organization{ID: organizationId}).Error
	return err
}

func (d *OrganizationDao) Update(organization model.Organization) error {
	err := d.db.Model(&organization).Updates(&organization).Error
	return err
}

func (d *OrganizationDao) Get(organizationId uint) (organization *model.Organization, error error) {
	var p model.Organization
	p.ID = organizationId

	err := d.db.Preload("Projects").Preload("Users").Find(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (d *OrganizationDao) UserOrganization(uid uint) (organizations []*model.Organization, error error) {
	u := model.User{}
	err := d.db.Debug().Where("id=?", uid).
		Preload("Organizations").
		Preload("Organizations.Projects").
		Preload("Organizations.Users").Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u.Organizations, nil
}

func (d *OrganizationDao) List(cur, size int) (result []model.Organization, current, pageSize int, total int64, error error) {

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
	list := make([]model.Organization, 0)

	err := d.db.Model(model.Organization{}).Preload("Projects").Preload("Users").
		Count(&t).Limit(s).Offset((n - 1) * s).Order("id desc").Find(&list).Error

	if err != nil {
		return list, n, s, t, err
	}
	return list, n, s, t, nil
}
