package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	NickName string     `gorm:"type:string;size:100;comment:昵称" json:"nickname"`
	Avatar   string     `gorm:"type:string;size:300;comment:头像" json:"avatar"`
	Email    string     `gorm:"type:string;size:100;comment:邮箱" json:"email"`
	Password string     `gorm:"type:string;size:500;comment:密码" json:"-"`
	Status   uint       `gorm:"comment:状态" json:"status"`
	RoleId   uint       `gorm:"comment:角色" json:"role_id"`
	Projects []*Project `gorm:"many2many:user_projects;comment:项目" json:"projects,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "user"
}
