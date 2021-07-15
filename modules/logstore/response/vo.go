package response

type LogList struct {
	Total int                      `json:"total"`
	Logs  []map[string]interface{} `json:"logs"`
}

// 日志
type LogsResponse struct {
	Count      int64                    `json:"count"`
	EffectUser int64                    `json:"effectUser"`
	Logs       []map[string]interface{} `json:"logs"`
}

// 通用统计列表
type CountListRes struct {
	Total int          `json:"total"`
	List  []*CountItem `json:"list"`
}

type CountItem struct {
	Key   string `json:"kye"`
	Count int64  `json:"count"`
	User  int64  `json:"user"`
}
