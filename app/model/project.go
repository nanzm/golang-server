package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId             string `gorm:"uniqueIndex; type:varchar(50);" json:"appId"`
	Name              string `gorm:"uniqueIndex; type:varchar(50);" json:"name"`
	Alias             string `gorm:"type:varchar(50);" json:"alias"`
	Type              string `gorm:"type:varchar(50);" json:"type"`
	GitRepositoryUrl  string `gorm:"type:varchar(500);" json:"git_repository_url"`
	GitRepositoryName string `gorm:"type:varchar(500);" json:"git_repository_name"`

	OrganizationId uint `gorm:"type:int(64);" json:"organization_id"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Project) TableName() string {
	return "project"
}
