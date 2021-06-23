package entity

import (
	"gorm.io/gorm"
	"time"
)

type Artifact struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Project  string `gorm:"type:string;size:100;comment:项目名" json:"project"`
	Link     string `gorm:"type:string;size:500;comment:文件链接" json:"link"`
	GitName  string `gorm:"type:string;size:300;comment:git用户名" json:"username"`
	GitEmail string `gorm:"type:string;size:20;comment:邮箱" json:"email"`
	GitRef   string `gorm:"type:string;size:100;comment:分支" json:"ref"`
	GitSha   string `gorm:"type:string;size:200;comment:sha" json:"sha"`
	GitMsg   string `gorm:"type:string;size:300;comment:commit" json:"commit"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Artifact) TableName() string {
	return "artifact"
}
