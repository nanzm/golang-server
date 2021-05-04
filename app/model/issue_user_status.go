package model

import (
	"time"

	"gorm.io/gorm"
)

type IssueUserStatus struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserId  uint `gorm:"type:int(64)" json:"user_count"`
	IssueId uint `gorm:"type:int(64)" json:"issue_id"`
	Read    bool `gorm:"type:int(64)" json:"read"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (IssueUserStatus) TableName() string {
	return "issue_user_status"
}
