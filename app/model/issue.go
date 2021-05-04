package model

import (
	"gorm.io/gorm"
	"time"
)

type Issue struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	AppId      string `gorm:"" json:"appId"`
	AppVersion string `gorm:"" json:"app_version"`
	Md5        string `gorm:"" json:"md5"`
	Env        string `gorm:"" json:"env"`
	Type       string `gorm:"" json:"type"`
	Category   string `gorm:"" json:"category"`
	Raw        string `gorm:"" json:"raw"`

	EventCount int  `gorm:"" json:"events_count"`
	UserCount  int  `gorm:"" json:"user_count"`
	Resolve    bool `gorm:"" json:"resolve"`
	Ignore     bool `gorm:"" json:"ignore"`
}

func (Issue) TableName() string {
	return "issue"
}
