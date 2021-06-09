package dto

type ListSearchParam struct {
	Current   int    `form:"current" binding:"number"`
	PageSize  int    `form:"pageSize" binding:"number"`
	SearchStr string `form:"search_str"`
}

type ListParam struct {
	Current  int `form:"current" binding:"number"`
	PageSize int `form:"pageSize" binding:"number"`
}
