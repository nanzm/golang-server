package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	NickName string     `gorm:"comment:昵称" json:"nickname"`
	Avatar   string     `gorm:"comment:头像" json:"avatar"`
	Email    string     `gorm:"comment:邮箱" json:"email"`
	Password string     `gorm:"comment:密码" json:"-"`
	Status   uint       `gorm:"comment:状态" json:"status"`
	RoleId   uint       `gorm:"comment:角色" json:"role_id"`
	Projects []*Project `gorm:"many2many:user_projects;comment:项目" json:"projects,omitempty"`
}

func (User) TableName() string {
	return "user"
}
