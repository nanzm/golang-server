package dto

import (
	"dora/internal/apps/manage/entity"
)

type SignUpParam struct {
	Email      string `json:"email" binding:"required,email"`
	NickName   string `json:"nickname" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmailLoginParam struct {
	Email   string `json:"email" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
}

type CaptchaParam struct {
	Type  string `json:"type" binding:"required"`
	Email string `json:"email" binding:"required"`
	//Phone string `json:"phone"`
}

type UpdateParam struct {
	Id         uint   `json:"id" binding:"required"`
	NickName   string `json:"nickname"`
	Email      string `json:"email" binding:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword" binding:"eqfield=Password"`
}

type StatusParam struct {
	Id     uint `json:"id" binding:"required"`
	Status uint `json:"status" binding:"required"`
}

type UserWithRole struct {
	entity.User
	Role *entity.Role `gorm:"foreignKey:RoleId;" json:"role"`
}

type UserLoginVo struct {
	Token string        `json:"token"`
	User  *UserWithRole `json:"user"`
}

type UpdateDefaultSettingReq struct {
	ProjectId      uint `json:"project_id"  binding:"required"`
}

type UserSettingVo struct {
	entity.UserSetting
	Project *entity.Project `gorm:"foreignKey:project_id;" json:"curProject"`
}
