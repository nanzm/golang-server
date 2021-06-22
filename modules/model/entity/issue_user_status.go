package entity

import (
	"gorm.io/gorm"
	"time"
)

type IssueUserStatus struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserId  uint `gorm:"" json:"user_count"`
	IssueId uint `gorm:"" json:"issue_id"`
	Read    bool `gorm:"" json:"read"`
}

func (IssueUserStatus) TableName() string {
	return "issue_user_status"
}
