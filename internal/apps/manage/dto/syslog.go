package dto

// nsq 消息解析失败
type ParseErrorParam struct {
	Current  int64 `form:"current" binding:"number"`
	PageSize int64 `form:"pageSize" binding:"number"`
}
