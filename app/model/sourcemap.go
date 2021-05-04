package model

import (
	"time"

	"gorm.io/gorm"
)

type SourceMap struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId      string `gorm:"type:int(64)" json:"appId"`
	AppVersion string `gorm:"type:int(64)" json:"app_version"`
	AppType    string `gorm:"type:int(64)" json:"app_type"`

	Version string `gorm:"type:int(64)" json:"version"`
	Path    string `gorm:"type:int(64)" json:"path"`
	Size    string `gorm:"type:int(64)" json:"size"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SourceMap) TableName() string {
	return "sourcemap"
}
