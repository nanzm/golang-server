package dto

import "mime/multipart"

type CreateProject struct {
	Name              string `json:"name" binding:"required"`
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

type SearchParam struct {
	Body map[string]interface{} `form:"body" binding:"required"`
}

//
type AlarmListParam struct {
	Current  int64  `form:"current" binding:"number"`
	PageSize int64  `form:"pageSize" binding:"number"`
	AppId    string `form:"appId"`
	//Start    int64  `form:"start" binding:"required"`
	//End      int64  `form:"end" binding:"required"`
}

// 备份
type BackUpParam struct {
	AppId       string `form:"appId" binding:"required"`
	ProjectName string `form:"project_name" binding:"required"`

	File     *multipart.FileHeader `form:"file" binding:"required"`
	FileName string                `form:"file_name" binding:"required"`
	FileType string                `form:"file_type" binding:"required"`

	GitName   string `form:"git_name"`
	GitEmail  string `form:"git_email"`
	GitBranch string `form:"git_branch"`

	Commit    string `form:"commit" binding:"required"`
	CommitSha string `form:"commit_sha" binding:"required"`
	CommitTs  string `form:"commit_ts" binding:"required"`
}

type BackUpListParam struct {
	Current  int64  `form:"current" binding:"number"`
	PageSize int64  `form:"pageSize" binding:"number"`
	FileType string `form:"file_type"`
	AppId    string `form:"appId"`
	//Start    int64  `form:"start" binding:"required"`
	//End      int64  `form:"end" binding:"required"`
}

type SourcemapParseParam struct {
	AppId string `json:"appId" binding:"required"`
	Stack string `json:"stack" binding:"required"`
}

type SourcemapDeleteParam struct {
	Id uint `json:"id" binding:"required"`
}

// 上传
type UploadSourcemapParam struct {
	AppId       string `form:"appId" binding:"required"`
	AppVersion  string `form:"appVersion"`
	ProjectName string `form:"project_name" binding:"required"`

	File     *multipart.FileHeader `form:"file" binding:"required"`
	FileName string                `form:"file_name" binding:"required"`
}

type SourcemapListParam struct {
	Current  int64  `form:"current" binding:"number"`
	PageSize int64  `form:"pageSize" binding:"number"`
	AppId    string `form:"appId"`
	//Start    int64  `form:"start" binding:"required"`
	//End      int64  `form:"end" binding:"required"`
}

// 访问 node 解析服务时的参数
type ReqPostParseData struct {
	Stack        string `json:"stack"`
	RawSourcemap string `json:"rawSourcemap"`
}
