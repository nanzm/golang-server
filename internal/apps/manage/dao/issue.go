package dao

import (
	"dora/internal/apps/manage/entity"
	dataGorm "dora/internal/datasource/gorm"
	"gorm.io/gorm"
)

type IssueDao struct {
	db *gorm.DB
}

func NewIssueDao() *IssueDao {
	return &IssueDao{
		db: dataGorm.Instance(),
	}
}

func (d *IssueDao) Create(issue *entity.Issue) (result *entity.Issue, error error) {
	err := d.db.Model(&entity.Issue{}).Create(issue).Error
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (d *IssueDao) Delete(issueId uint) error {
	err := d.db.Delete(&entity.Issue{ID: issueId}).Error
	return err
}

func (d *IssueDao) Update(issue *entity.Issue) error {
	err := d.db.Model(&issue).Updates(issue).Error
	return err
}

func (d *IssueDao) UpdateByMd5(md5 string, issue *entity.Issue) error {
	err := d.db.Model(&issue).Where("md5 = ? ", md5).Updates(issue).Error
	return err
}

func (d *IssueDao) QueryByMd5(md5 string) (issue *entity.Issue, error error) {
	var p entity.Issue
	error = d.db.Model(&entity.Issue{}).Where("md5 = ? ", md5).Find(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *IssueDao) Get(issueId uint) (issue *entity.Issue, error error) {
	var p entity.Issue
	p.ID = issueId

	error = d.db.Where("id = ?", issueId).Find(&p).Error
	if error != nil {
		return nil, error
	}
	return &p, nil
}

func (d *IssueDao) ListQueryTimeRange(appId string, start, end int64, cur, size int) (
	result []*entity.Issue, current, pageSize int, total int64, error error) {

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
	list := make([]*entity.Issue, 0)

	err := d.db.Debug().Model(entity.Issue{}).
		Where("app_id = ?", appId).
		Where("updated_at BETWEEN to_timestamp(?) AND to_timestamp(?)", start, end).
		Count(&t).Limit(s).Offset((n - 1) * s).
		Order("updated_at desc").Find(&list).Error

	if err != nil {
		return list, n, s, t, err
	}
	return list, n, s, t, nil
}
