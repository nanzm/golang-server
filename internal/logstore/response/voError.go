package response


// 错误
type ErrorCountRes struct {
	Count      int `json:"count"`
	EffectUser int `json:"effect_user"`
}

// 错误列表
type ErrorListRes struct {
	Total int          `json:"total"`
	List  []*ErrorItem `json:"list"`
}

type ErrorItem struct {
	Md5        string `json:"md5"`
	Msg        string `json:"msg"`
	Error      string `json:"error"`
	Count      int    `json:"count"`
	EffectUser int    `json:"effectUser"`
	FirstAt    int    `json:"firstAt"`
	LastAt     int    `json:"lastAt"`
}
