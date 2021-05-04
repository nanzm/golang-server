package model

import (
	"time"

	"gorm.io/gorm"
)

type Issue struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId      string `gorm:"type:varchar(50)" json:"appId"`
	AppVersion string `gorm:"type:varchar(200)" json:"app_version"`
	Md5        string `gorm:"type:varchar(100)" json:"md5"`
	Env        string `gorm:"type:varchar(30)" json:"env"`
	Type       string `gorm:"type:varchar(30)" json:"type"`
	Category   string `gorm:"type:varchar(30)" json:"category"`
	Raw        string `gorm:"type:text" json:"raw"`

	EventCount int  `gorm:"type:int(64)" json:"events_count"`
	UserCount  int  `gorm:"type:int(64)" json:"user_count"`
	Resolve    bool `gorm:"type:boolean" json:"resolve"`
	Ignore     bool `gorm:"type:boolean" json:"ignore"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Issue) TableName() string {
	return "issue"
}
