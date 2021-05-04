package dto

import "mime/multipart"

type CreateProject struct {
	Name              string `json:"name" binding:"required"`
	OrganizationId    uint   `json:"organization_id" binding:"required"`
	Type              string `json:"type" binding:"required"`
	Alias             string `json:"alias"`
	GitRepositoryUrl  string `json:"git_repository_url"`
	GitRepositoryName string `json:"git_repository_name"`
}

type QueryDetail struct {
	ProjectId uint `form:"project_id"  binding:"required"`
}

type ReqOrganizationProjectsList struct {
	OrganizationId uint `form:"organizationId"  binding:"required"`
}

// 上传
type UploadSourcemapParam struct {
	AppId string                `form:"appId" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

// 备份
type BackUpParam struct {
	AppId       string                `form:"appId" binding:"required"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
	Commit      string                `form:"commit" binding:"required"`
	ProjectName string                `form:"projectName" binding:"required"`
}

// 管理平台调用
type SourcemapParseParam struct {
	AppId string `json:"appId" binding:"required"`
	Stack string `json:"stack" binding:"required"`
}

// 访问 node 解析服务时的参数
type ReqPostParseData struct {
	Stack        string `json:"stack"`
	RawSourcemap string `json:"rawSourcemap"`
}
