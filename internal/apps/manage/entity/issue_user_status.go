package entity

import (
	"gorm.io/gorm"
	"time"
)

type IssueUserStatus struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserId  uint `gorm:"type:int;" json:"user_count"`
	IssueId uint `gorm:"type:int;" json:"issue_id"`
	Read    bool `gorm:"type:bool;" json:"read"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (IssueUserStatus) TableName() string {
	return "issue_user_status"
}
