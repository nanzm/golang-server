package model

import (
	"time"

	"gorm.io/gorm"
)

type SysLog struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Error   string `gorm:"type:text" json:"error"`
	RawData string `gorm:"type:text" json:"raw_data"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SysLog) TableName() string {
	return "syslog"
}
