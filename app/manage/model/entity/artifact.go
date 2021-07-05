package entity

import (
	"gorm.io/gorm"
	"time"
)

type Artifact struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId       string `gorm:"type:string;size:100;" json:"appId"`
	ProjectName string `gorm:"type:string;size:100;comment:项目名" json:"project_name"`

	FileName string `gorm:"type:string;size:300;comment:文件名" json:"file_name"`
	FileType string `gorm:"type:string;size:50;comment:文件类型" json:"file_type"`
	FilePath string `gorm:"type:string;size:300;comment:文件路径" json:"file_path"`

	GitName   string `gorm:"type:string;size:30;comment:git用户名" json:"git_username"`
	GitEmail  string `gorm:"type:string;size:30;comment:git邮箱" json:"git_email"`
	GitBranch string `gorm:"type:string;size:30;comment:git分支" json:"git_branch"`

	Commit    string `gorm:"type:string;size:200;comment:git sha" json:"commit"`
	CommitSha string `gorm:"type:string;size:300;comment:git commit" json:"commit_sha"`
	CommitTs  string `gorm:"type:string;size:300;comment:git commit" json:"commit_ts"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Artifact) TableName() string {
	return "artifact"
}
