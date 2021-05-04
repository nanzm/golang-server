package model

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Name         string `gorm:"type:varchar(100)" json:"name"`
	Introduction string `gorm:"type:varchar(255)" json:"introduction"`
	Type         string `gorm:"type:varchar(255)" json:"type"`
	CreateUid    uint   `gorm:"type:int(64)" json:"create_uid"`

	Users    []User    `gorm:"many2many:user_organizations;" json:"users"`
	Projects []Project `gorm:"foreignKey:OrganizationId" json:"projects"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Organization) TableName() string {
	return "organization"
}
