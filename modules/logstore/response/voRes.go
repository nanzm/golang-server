package response

type ResLoadFailTotalRes struct {
	Count int `json:"count"`
	User  int `json:"user"`
}

// 资源加载时间
type ResDurationRes struct {
	List    []*ResDurationItemRes `json:"list"`
	Percent *ResDurationPercent   `json:"percent"`
}

type ResDurationPercent struct {
	Percent1  float64 `json:"p1"`
	Percent5  float64 `json:"p5"`
	Percent25 float64 `json:"p25"`
	Percent50 float64 `json:"p50"`
	Percent75 float64 `json:"p75"`
	Percent95 float64 `json:"p95"`
	Percent99 float64 `json:"p99"`
}

type ResDurationItemRes struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}


type ResTopListDurationRes struct {
	Total int           `json:"total"`
	List  []*ResTopItem `json:"list"`
}

type ResTopItem struct {
	Key       string  `json:"key"`
	Count     int64   `json:"count"`
	User      int64   `json:"user"`
	Percent1  float64 `json:"p1"`
	Percent5  float64 `json:"p5"`
	Percent25 float64 `json:"p25"`
	Percent50 float64 `json:"p50"`
	Percent75 float64 `json:"p75"`
	Percent95 float64 `json:"p95"`
	Percent99 float64 `json:"p99"`
}
