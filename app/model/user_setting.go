package model

import (
	"time"

	"gorm.io/gorm"
)

type UserSetting struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserId         uint `gorm:"type:int(64);uniqueIndex;" json:"user_id"`
	OrganizationId uint `gorm:"type:int(64);not null;" json:"organization_id"`
	ProjectId      uint `gorm:"type:int(64);" json:"project_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (UserSetting) TableName() string {
	return "user_setting"
}
