package response

type PerfMetricsBucket struct {
	Fp   []*PerfMetricsBucketItem `json:"fp"`
	Fcp  []*PerfMetricsBucketItem `json:"fcp"`
	Lcp  []*PerfMetricsBucketItem `json:"lcp"`
	Fid  []*PerfMetricsBucketItem `json:"fid"`
	Cls  []*PerfMetricsBucketItem `json:"cls"`
	Ttfb []*PerfMetricsBucketItem `json:"ttfb"`
}

type PerfMetricsBucketItem struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}
