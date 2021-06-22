package entity

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	AppId             string `gorm:"uniqueIndex;" json:"appId"`
	Name              string `gorm:"uniqueIndex;" json:"name"`
	Alias             string `gorm:"" json:"alias"`
	Type              string `gorm:"" json:"type"`
	GitRepositoryUrl  string `gorm:"" json:"git_repository_url"`
	GitRepositoryName string `gorm:"" json:"git_repository_name"`

	Users []*User `gorm:"many2many:user_projects;" json:"users,omitempty"`
}

func (Project) TableName() string {
	return "project"
}
