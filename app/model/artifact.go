package model

import (
	"time"

	"gorm.io/gorm"
)

type Artifact struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Project  string `gorm:"type:varchar(255);comment:项目名" json:"project"`
	Name     string `gorm:"type:varchar(255);comment:文件名" json:"name"`
	Link     string `gorm:"type:varchar(500);comment:文件链接" json:"link"`
	GitName  string `gorm:"type:varchar(255);comment:git用户名" json:"username"`
	GitEmail string `gorm:"type:varchar(255);comment:邮箱" json:"email"`
	GitRef   string `gorm:"type:varchar(255);comment:分支" json:"ref"`
	GitSha   string `gorm:"type:varchar(255);comment:sha" json:"sha"`
	GitMsg   string `gorm:"type:varchar(500);comment:commit" json:"commit"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Artifact) TableName() string {
	return "artifact"
}
