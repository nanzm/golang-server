package dto

type ErrorParam struct {
	Status int `form:"status" json:"status" binding:"required"`
}

type DelayParam struct {
	Delay int64 `form:"delay" json:"delay" binding:"required"`
}
