package entity

import (
	"gorm.io/gorm"
	"time"
)

type SysLog struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Error   string `gorm:"type:string;size:500;" json:"error"`
	RawData string `gorm:"type:TEXT;" json:"raw_data"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SysLog) TableName() string {
	return "syslog"
}
