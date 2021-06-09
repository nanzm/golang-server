package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Key     string `gorm:"uniqueIndex;comment:标识" json:"key"`
	Name    string `gorm:"comment:角色名" json:"name"`
	Remarks string `gorm:"comment:备注" json:"remarks"`

	Users []*User `gorm:"foreignKey:RoleId;references:ID" json:"users,omitempty"`
}

func (Role) TableName() string {
	return "role"
}
