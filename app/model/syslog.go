package model

import (
	"gorm.io/gorm"
	"time"
)

type SysLog struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Error   string `gorm:"" json:"error"`
	RawData string `gorm:"" json:"raw_data"`
}

func (SysLog) TableName() string {
	return "syslog"
}
