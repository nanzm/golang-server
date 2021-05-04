package model

import (
	"gorm.io/gorm"
	"time"
)

type Artifact struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Project  string `gorm:"comment:项目名" json:"project"`
	Name     string `gorm:"comment:文件名" json:"name"`
	Link     string `gorm:"comment:文件链接" json:"link"`
	GitName  string `gorm:"comment:git用户名" json:"username"`
	GitEmail string `gorm:"comment:邮箱" json:"email"`
	GitRef   string `gorm:"comment:分支" json:"ref"`
	GitSha   string `gorm:"comment:sha" json:"sha"`
	GitMsg   string `gorm:"comment:commit" json:"commit"`
}

func (Artifact) TableName() string {
	return "artifact"
}
