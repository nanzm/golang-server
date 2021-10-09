package response

// 接口错误
type ApiErrorCountRes struct {
	Count      int `json:"count"`
	EffectUser int `json:"effect_user"`
}

type ApiErrorTrendRes struct {
	Total int                     `json:"total"`
	List  []*ApiErrorTrendItemRes `json:"list"`
}

type ApiErrorTrendItemRes struct {
	Count      int    `json:"count"`
	EffectUser int    `json:"effect_user"`
	Ts         string `json:"ts"`
}

// 接口加载时间趋势
type ApiDurationTrendRes struct {
	Total int                        `json:"total"`
	List  []*ApiDurationTrendItemRes `json:"list"`
}

type ApiDurationTrendItemRes struct {
	Percent1  float64 `json:"p1"`
	Percent5  float64 `json:"p5"`
	Percent25 float64 `json:"p25"`
	Percent50 float64 `json:"p50"`
	Percent75 float64 `json:"p75"`
	Percent95 float64 `json:"p95"`
	Percent99 float64 `json:"p99"`
	Ts        string  `json:"ts"`
}

// 接口时间加载
type ApiDurationRes struct {
	List    []*ApiDurationItemRes `json:"list"`
	Percent *ApiDurationPercent   `json:"percent"`
}

type ApiDurationItemRes struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}

type ApiDurationPercent struct {
	Percent1  float64 `json:"p1"`
	Percent5  float64 `json:"p5"`
	Percent25 float64 `json:"p25"`
	Percent50 float64 `json:"p50"`
	Percent75 float64 `json:"p75"`
	Percent95 float64 `json:"p95"`
	Percent99 float64 `json:"p99"`
}

type ApiTopListDurationRes struct {
	Total int           `json:"total"`
	List  []*ApiTopItem `json:"list"`
}

type ApiTopItem struct {
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

// 接口错误
type ApiErrorListRes struct {
	Total int             `json:"total"`
	List  []*ApiErrorItem `json:"list"`
}

type ApiErrorItem struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Method     string `json:"method"`
	ErrorType  string `json:"error_type"`
	Count      int    `json:"count"`
	EffectUser int    `json:"effect_user"`
}
