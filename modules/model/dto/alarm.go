package dto

import (
	"dora/modules/model/entity"
)

type AlarmProject struct {
	ProjectInfo entity.Project `gorm:"foreignKey:ProjectId" json:"project_info"`
}

type AlarmLog struct {
	entity.AlarmLog
}
