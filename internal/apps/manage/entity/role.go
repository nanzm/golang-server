package entity

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Key     string `gorm:"type:string;size:40;uniqueIndex;comment:标识;" json:"key"`
	Name    string `gorm:"type:string;size:100;comment:角色名;" json:"name"`
	Remarks string `gorm:"type:string;size:300;comment:备注;" json:"remarks"`

	Users []*User `gorm:"foreignKey:RoleId;references:ID" json:"users,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Role) TableName() string {
	return "role"
}
