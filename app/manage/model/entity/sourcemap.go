package entity

import (
	"gorm.io/gorm"
	"time"
)

type Sourcemap struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId      string `gorm:"type:string;size:100" json:"appId"`
	AppVersion string `gorm:"type:string;size:300" json:"app_version"`

	Path string `gorm:"type:string;size:300" json:"path"`
	Size int64  `gorm:"type:int;" json:"size"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Sourcemap) TableName() string {
	return "sourcemap"
}
