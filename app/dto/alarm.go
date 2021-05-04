package dto

import "dora/app/model"

type AlarmProject struct {
	ProjectInfo model.Project `gorm:"foreignKey:ProjectId" json:"project_info"`
}

type AlarmLog struct {
	model.AlarmLog
}
