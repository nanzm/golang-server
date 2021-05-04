package model

import (
	"time"

	"gorm.io/gorm"
)

type AlarmLog struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AlarmProjectId uint   `gorm:"type:int(64);comment:告警id" json:"alarm_project_id"`
	Log            string `gorm:"type:varchar(255);comment:日志" json:"log"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 告警对象
func (AlarmLog) TableName() string {
	return "alarm_log"
}
