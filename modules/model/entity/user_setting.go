package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserSetting struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserId    uint `gorm:"uniqueIndex;" json:"user_id"`
	ProjectId uint `gorm:"type:int;" json:"project_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (UserSetting) TableName() string {
	return "user_setting"
}
