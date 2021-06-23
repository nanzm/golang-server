package dto

// nsq 消息解析失败
type ParseErrorParam struct {
	Current  int `form:"current" binding:"number"`
	PageSize int `form:"pageSize" binding:"number"`
}
