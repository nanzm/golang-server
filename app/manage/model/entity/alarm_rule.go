package entity

import (
	"gorm.io/gorm"
	"time"
)

type AlarmStrategy struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId string `gorm:"type:string;size:100" json:"appId"`
	Level int    `gorm:"type:string;size:100;comment:级别" json:"period"`

	HookStatus  int    `gorm:"type:int;comment:状态" json:"hook_status"`
	HookType    string `gorm:"type:string;comment:类型" json:"hook_type"`
	HookAddress string `gorm:"type:string;comment:链接" json:"hook_address"`
	HookSign    string `gorm:"type:string;comment:加签" json:"hook_sign"`

	Expression string `gorm:"type:string;comment:表达式" json:"expression"`
	Content    string `gorm:"type:string;comment:告警内容" json:"content"`
	Status     int    `gorm:"comment:0:disable, 1:enable" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (AlarmStrategy) TableName() string {
	return "alarm_strategy"
}
