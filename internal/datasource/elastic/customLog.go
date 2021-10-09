package elastic

import (
	"dora/pkg/utils/logx"
	"net/http"
	"net/url"
	"time"
)

type customLog struct {
}

func (l *customLog) LogRoundTrip(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) error {
	query, _ := url.QueryUnescape(req.URL.RawQuery)
	if query != "" {
		query = "?" + query
	}

	var (
		status string
		color  string
	)

	status = res.Status
	switch {
	case res.StatusCode > 0 && res.StatusCode < 300:
		color = "\x1b[32m"
	case res.StatusCode > 299 && res.StatusCode < 500:
		color = "\x1b[33m"
	case res.StatusCode > 499:
		color = "\x1b[31m"
	default:
		status = "ERROR"
		color = "\x1b[31;4m"
	}

	if res.StatusCode > 299 {
		logx.Infof("%6s \x1b[1;4m%s://%s%s\x1b[0m%s %s%s\x1b[0m \x1b[2m%s\x1b[0m",
			req.Method,
			req.URL.Scheme,
			req.URL.Host,
			req.URL.Path,
			query,
			color,
			status,
			dur.Truncate(time.Millisecond),
		)
	}

	if err != nil {
		logx.Infof("\x1b[31;1m» ERROR \x1b[31m%v\x1b[0m", err)
	}

	return nil
}

func (l *customLog) RequestBodyEnabled() bool { return false }

func (l *customLog) ResponseBodyEnabled() bool { return false }
