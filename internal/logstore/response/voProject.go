package response

type PageUrlVisitListRes struct {
	Total int64               `json:"total"`
	List  []*PageUrlVisitItem `json:"list"`
}

type PageUrlVisitItem struct {
	Url string `json:"url"`
	Pv  int64  `json:"pv"`
	Uv  int64  `json:"uv"`
	Bu  int64  `json:"bu"`
}

// PV UV
type PvUvTotalRes struct {
	Pv int64 `json:"pv"`
	Uv int64 `json:"uv"`
}

type PvUvTrendRes struct {
	Total int                 `json:"total"`
	List  []*PvUvTrendItemRes `json:"list"`
}

type PvUvTrendItemRes struct {
	Pv int64  `json:"pv"`
	Uv int64  `json:"uv"`
	Ts string `json:"ts"`
}

// 项目事件数
type ProjectLogCountRes struct {
	Count int64 `json:"count"`
	User  int64 `json:"user"`
	Bu    int64 `json:"bu"`
}
