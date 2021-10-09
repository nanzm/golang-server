package entity

import (
	"gorm.io/gorm"
	"time"
)

type AlarmContact struct {
	ID uint `gorm:"primaryKey" json:"id"`

	AlarmId uint   `gorm:"comment:告警id" json:"alarm_id"`
	Type    string `gorm:"type:string;size:10;comment:联系人 user、email、dingTalk" json:"type"`

	Remark string `gorm:"type:string;size:300;comment:备注" json:"remark"`

	// 联系人
	UserId uint `gorm:"comment:联系人" json:"user_id"`
	// 钉钉
	DingTalkAT     string `gorm:"type:string;size:100;comment:钉钉access_token" json:"ding_talk_at"`
	DingTalkSecret string `gorm:"type:string;size:100;comment:钉钉secret" json:"ding_talk_secret"`

	// email
	Email string `gorm:"type:string;size:50;comment:邮箱" json:"email"`

	Status int `gorm:"comment:0:开启, 1:关闭" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (AlarmContact) TableName() string {
	return "alarm_contact"
}
