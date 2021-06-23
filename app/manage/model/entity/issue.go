package entity

import (
	"gorm.io/gorm"
	"time"
)

type Issue struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId      string `gorm:"type:string;size:100" json:"appId"`
	AppVersion string `gorm:"type:string;size:100" json:"app_version"`
	Md5        string `gorm:"type:string;size:100" json:"md5"`
	Env        string `gorm:"type:string;size:100" json:"env"`
	Type       string `gorm:"type:string;size:100" json:"type"`
	Category   string `gorm:"type:string;size:100" json:"category"`
	Raw        string `gorm:"type:string;" json:"raw"`

	EventCount int  `gorm:"type:int;" json:"events_count"`
	UserCount  int  `gorm:"type:int;" json:"user_count"`
	Resolve    bool `gorm:"type:bool;" json:"resolve"`
	Ignore     bool `gorm:"type:bool;" json:"ignore"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Issue) TableName() string {
	return "issue"
}
