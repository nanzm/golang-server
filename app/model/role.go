package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Key     string `gorm:"type:varchar(50); uniqueIndex;" json:"key"`
	Name    string `gorm:"type:varchar(50)" json:"name"`
	Remarks string `gorm:"type:varchar(50)" json:"remarks"`

	Users []User `gorm:"foreignKey:RoleId;references:ID" json:"users"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Role) TableName() string {
	return "role"
}
