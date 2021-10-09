package entity

import (
	"gorm.io/gorm"
	"time"
)

type Alarm struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AppId string `gorm:"type:string;size:100" json:"appId"`

	Level    string `gorm:"type:string;size:100;comment:级别" json:"level"`
	RuleType string `gorm:"type:string;size:100;comment:类型" json:"rule"`

	Time     int    `gorm:"type:int;comment:时间" json:"time"`
	TimeUnit string `gorm:"type:string;size:10;comment:时间单位" json:"timeUnit"`
	Operator string `gorm:"type:string;size:10;comment:比较符号" json:"operator"`
	Quota    int    `gorm:"type:int;comment:值" json:"quota"`

	Content string `gorm:"type:string;size:500;comment:告警内容" json:"content"`

	Status int `gorm:"comment:0:disable, 1:enable" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Alarm) TableName() string {
	return "alarm"
}
