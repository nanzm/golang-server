package entity

import (
	"gorm.io/gorm"
	"time"
)

type SourceMap struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId      string `gorm:"type:string;" json:"appId"`
	AppVersion string `gorm:"type:string;" json:"app_version"`
	AppType    string `gorm:"type:string;" json:"app_type"`

	Version string `gorm:"type:string;" json:"version"`
	Path    string `gorm:"type:string;" json:"path"`
	Size    string `gorm:"type:string;" json:"size"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (SourceMap) TableName() string {
	return "sourcemap"
}
