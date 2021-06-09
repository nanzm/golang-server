package model

import (
	"gorm.io/gorm"
	"time"
)

type SourceMap struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	AppId      string `gorm:"" json:"appId"`
	AppVersion string `gorm:"" json:"app_version"`
	AppType    string `gorm:"" json:"app_type"`

	Version string `gorm:"" json:"version"`
	Path    string `gorm:"" json:"path"`
	Size    string `gorm:"" json:"size"`
}

func (SourceMap) TableName() string {
	return "sourcemap"
}
