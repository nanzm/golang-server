package entity

import (
	"gorm.io/gorm"
	"time"
)

type AlarmLog struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AlarmId        uint   `gorm:"comment:告警id" json:"alarm_id"`
	AlarmContactId uint   `gorm:"comment:告警联系方式id" json:"alarm_contact_id"`
	Content        string `gorm:"type:string;size:500;comment:内容" json:"content"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 告警对象
func (AlarmLog) TableName() string {
	return "alarm_log"
}
