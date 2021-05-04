package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey" json:"id"`

	NickName string `gorm:"type:varchar(50);comment:昵称" json:"nickname"`
	Avatar   string `gorm:"type:varchar(255);comment:头像" json:"avatar"`
	Email    string `gorm:"type:varchar(255);comment:邮箱" json:"email"`
	Password string `gorm:"type:varchar(255);comment:密码" json:"-"`
	Status   uint   `gorm:"type:int(10);comment:状态" json:"status"`
	RoleId   uint   `gorm:"type:int(64)" json:"role_id"`

	Organizations []*Organization `gorm:"many2many:user_organizations;" json:"organizations"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "user"
}
