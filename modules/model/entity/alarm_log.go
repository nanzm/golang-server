package entity

import (
	"gorm.io/gorm"
	"time"
)

type AlarmLog struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AlarmProjectId uint   `gorm:"comment:告警id" json:"alarm_project_id"`
	Log            string `gorm:"type:text;comment:日志" json:"log"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 告警对象
func (AlarmLog) TableName() string {
	return "alarm_log"
}
