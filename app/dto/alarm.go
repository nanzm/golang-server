package dto

import "dora/app/model"

type AlarmProject struct {
	model.AlarmProject
	ProjectInfo model.Project `gorm:"foreignKey:ProjectId" json:"project_info"`
}

type AlarmLog struct {
	model.AlarmLog
	AlarmProject model.AlarmProject `gorm:"foreignKey:AlarmProjectId" json:"project_info"`
}
