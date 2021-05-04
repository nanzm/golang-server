package model

import (
	"gorm.io/gorm"
	"time"
)

type AlarmStrategy struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	AppId string `gorm:"" json:"appId"`
	Level int    `gorm:"comment:级别" json:"period"`

	HookStatus  int    `gorm:"comment:状态" json:"hook_status"`
	HookType    string `gorm:"comment:类型" json:"hook_type"`
	HookAddress string `gorm:"comment:链接" json:"hook_address"`
	HookSign    string `gorm:"comment:加签" json:"hook_sign"`

	Expression string `gorm:"comment:表达式" json:"expression"`
	Content    string `gorm:"comment:告警内容" json:"content"`
	Status     int    `gorm:"comment:0:disable, 1:enable" json:"status"`
}

func (AlarmStrategy) TableName() string {
	return "alarm_strategy"
}
