package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserSetting struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserId    uint `gorm:"uniqueIndex;" json:"user_id"`
	ProjectId uint `gorm:"" json:"project_id"`
}

func (UserSetting) TableName() string {
	return "user_setting"
}
