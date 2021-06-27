package response

type LogsResponse struct {
	Count int64               `json:"count"`
	Logs  []map[string]string `json:"logs"`
}

// 统计 md 的次数
type LogCountByMd5Res struct {
	Count      int `json:"count"`
	EffectUser int `json:"effect_user"`
}

// PV UV
type PvUvTotalRes struct {
	Pv int `json:"pv"`
	Uv int `json:"uv"`
}

type PvUvTrendRes struct {
	Total int                 `json:"total"`
	List  []*PvUvTrendItemRes `json:"list"`
}

type PvUvTrendItemRes struct {
	Pv int    `json:"pv"`
	Uv int    `json:"uv"`
	Ts string `json:"ts"`
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

// 统计 md 的次数
type ErrorListRes struct {
	Total int          `json:"total"`
	List  []*ErrorItem `json:"list"`
}

// sdk 版本
type SdkVersionCountRes struct {
	Total int               `json:"total"`
	List  []*SdkVersionItem `json:"list"`
}

type SdkVersionItem struct {
	Version string `json:"version"`
	Count   int    `json:"count"`
}

// 上报类目
type CategoryCountRes struct {
	Total int                  `json:"total"`
	List  []*CategoryCountItem `json:"list"`
}

type CategoryCountItem struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// 页面统计
type PageTotalRes struct {
	Total int                 `json:"total"`
	List  []*PageTotalItemRes `json:"list"`
}

type PageTotalItemRes struct {
	Url string `json:"url"`
	Pv  int    `json:"pv"`
	Uv  int    `json:"uv"`
}

// 错误
type ErrorCountRes struct {
	Count      int `json:"count"`
	EffectUser int `json:"effect_user"`
}

type ErrorCountTrendRes struct {
	Total int                       `json:"total"`
	List  []*ErrorCountTrendItemRes `json:"list"`
}

type ErrorCountTrendItemRes struct {
	Count      int    `json:"count"`
	EffectUser int    `json:"effect_user"`
	Ts         string `json:"ts"`
}

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

// 网络
type PerfNavigationTimingTrendRes struct {
	Total int                             `json:"total"`
	List  []*NavigationTimingTrendItemRes `json:"list"`
}

type NavigationTimingTrendItemRes struct {
	DnsLookupTime   float64 `json:"dnsLookupTime"`
	DownloadTime    float64 `json:"downloadTime"`
	FetchTime       float64 `json:"fetchTime"`
	HeaderSize      float64 `json:"headerSize"`
	TimeToFirstByte float64 `json:"timeToFirstByte"`
	TotalTime       float64 `json:"totalTime"`
	Ts              string  `json:"ts"`
}

type PerfNavigationTimingValuesRes struct {
	DnsLookupTime   float64 `json:"dnsLookupTime"`
	DownloadTime    float64 `json:"downloadTime"`
	FetchTime       float64 `json:"fetchTime"`
	HeaderSize      float64 `json:"headerSize"`
	TimeToFirstByte float64 `json:"timeToFirstByte"`
	TotalTime       float64 `json:"totalTime"`
}

// 资源加载
type PerfDataConsumptionTrendRes struct {
	Total int                                `json:"total"`
	List  []*PerfDataConsumptionTrendItemRes `json:"list"`
}

type PerfDataConsumptionTrendItemRes struct {
	Css    float64 `json:"css"`
	Img    float64 `json:"img"`
	Other  float64 `json:"other"`
	Script float64 `json:"script"`
	Total  float64 `json:"total"`
	Xhr    float64 `json:"xhr"`
	Fetch  float64 `json:"fetch"`
	Ts     string  `json:"ts"`
}

type PerfDataConsumptionValuesRes struct {
	Css    float64 `json:"css"`
	Img    float64 `json:"img"`
	Other  float64 `json:"other"`
	Script float64 `json:"script"`
	Total  float64 `json:"total"`
	Xhr    float64 `json:"xhr"`
	Fetch  float64 `json:"fetch"`
}

// 性能
type PerfMetricsTrend struct {
	Total int                        `json:"total"`
	List  []*PerfMetricsTrendItemRes `json:"list"`
}

type PerfMetricsBucketItem struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

type PerfMetricsBucket struct {
	Fp   []*PerfMetricsBucketItem `json:"fp"`
	Fcp  []*PerfMetricsBucketItem `json:"fcp"`
	Lcp  []*PerfMetricsBucketItem `json:"lcp"`
	Fid  []*PerfMetricsBucketItem `json:"fid"`
	Cls  []*PerfMetricsBucketItem `json:"cls"`
	Ttfb []*PerfMetricsBucketItem `json:"ttfb"`
}

type PerfMetricsTrendItemRes struct {
	Fp  float64 `json:"fp"`
	Fcp float64 `json:"fcp"`
	Lcp float64 `json:"lcp"`
	Fid float64 `json:"fid"`
	Cls float64 `json:"cls"`
	Tbt float64 `json:"tbt"`
	Ts  string  `json:"ts"`
}

type PerfMetricsValuesRes struct {
	Fp  float64 `json:"fp"`
	Fcp float64 `json:"fcp"`
	Lcp float64 `json:"lcp"`
	Fid float64 `json:"fid"`
	Cls float64 `json:"cls"`
	Tbt float64 `json:"tbt"`
}

// 资源加载失败
type ResLoadFailTotalTrendRes struct {
	Total int                             `json:"total"`
	List  []*ResLoadFailTotalTrendItemRes `json:"list"`
}

type ResLoadFailTotalTrendItemRes struct {
	Count      int    `json:"count"`
	EffectUser int    `json:"effect_user"`
	Ts         string `json:"ts"`
}

type ResLoadFailTotalRes struct {
	Count      int `json:"count"`
	EffectUser int `json:"effect_user"`
}

type ResLoadFailListRes struct {
	Total int                   `json:"total"`
	List  []*ResLoadFailItemRes `json:"list"`
}

type ResLoadFailItemRes struct {
	Src        string `json:"src"`
	Count      int    `json:"count"`
	EffectUser int    `json:"effect_user"`
}

// Ip  国家 省份 城市
type ProjectIpToCountryRes struct {
	Country string `json:"country"`
	Pv      int    `json:"pv"`
	Uv      int    `json:"uv"`
}

type ProjectIpToProvinceRes struct {
	Province string `json:"province"`
	Pv       int    `json:"pv"`
	Uv       int    `json:"uv"`
}

type ProjectIpToCityRes struct {
	City string `json:"city"`
	Pv   int    `json:"pv"`
	Uv   int    `json:"uv"`
}

// 事件数
type ProjectEventCountRes struct {
	Count int `json:"count"`
}

// 发送模式
type ProjectSendModeRes struct {
	Total int             `json:"total"`
	List  []*SendModeItem `json:"list"`
}

type SendModeItem struct {
	Mode  string `json:"mode"`
	Count int    `json:"count"`
}

// 版本信息
type ProjectVersionRes struct {
	Total int            `json:"total"`
	List  []*VersionItem `json:"list"`
}

type VersionItem struct {
	Version string `json:"version"`
	Count   int    `json:"count"`
}

// 用户屏幕
type ProjectUserScreenRes struct {
	Total int           `json:"total"`
	List  []*ScreenItem `json:"list"`
}

type ScreenItem struct {
	ScreenWH string `json:"wh"`
	Count    int    `json:"count"`
}

// 事件类型
type ProjectCategoryRes struct {
	Total int             `json:"total"`
	List  []*CategoryItem `json:"list"`
}

type CategoryItem struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// 环境
type ProjectEnvRes struct {
	Total int        `json:"total"`
	List  []*EnvItem `json:"list"`
}

type EnvItem struct {
	Env   string `json:"env"`
	Count int    `json:"count"`
}
