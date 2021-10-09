package entity

import (
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID        uint           `gorm:"primaryKey" json:"id"`

	AppId             string `gorm:"type:string;uniqueIndex;size:100;" json:"appId"`
	Name              string `gorm:"type:string;uniqueIndex;size:100;" json:"name"`
	Alias             string `gorm:"type:string;size:300;" json:"alias"`
	Type              string `gorm:"type:string;size:10" json:"type"`
	GitRepositoryUrl  string `gorm:"type:string;size:500" json:"git_repository_url"`
	GitRepositoryName string `gorm:"type:string;size:300" json:"git_repository_name"`

	Users []*User `gorm:"many2many:user_projects;" json:"users,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Project) TableName() string {
	return "project"
}
