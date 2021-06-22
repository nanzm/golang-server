package dto

type IssueDetailParam struct {
	AppId string `form:"appId" binding:"required"`
	Md5   string `form:"md5" binding:"required"`
	Start int64  `form:"start" binding:"required"`
	End   int64  `form:"end" binding:"required"`
}

type IssueTrendParam struct {
	Md5 string `form:"md5" binding:"required"`
}

type QueryEventsByMd5Param struct {
	AppId string `form:"appId" binding:"required"`
	Md5   string `form:"md5" binding:"required"`
	Start int64  `form:"start" binding:"required"`
	End   int64  `form:"end" binding:"required"`
}

type IssueListParam struct {
	Current  int    `form:"current" binding:"number"`
	PageSize int    `form:"pageSize" binding:"number"`
	AppId    string `form:"appId" binding:"required"`
	Start    int64  `form:"start" binding:"required"`
	End      int64  `form:"end" binding:"required"`
}

type ChartData struct {
	AppId string `form:"appId" binding:"required"`
	Start int64  `form:"start" binding:"required,number"`
	End   int64  `form:"end" binding:"required,number"`
	// 间隔 单位分钟
	Interval int64 `form:"interval,default=60" binding:"number"`
}
