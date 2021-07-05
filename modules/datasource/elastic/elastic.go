package elastic

import (
	"dora/config"
	"dora/pkg/utils/logx"
	"github.com/cenkalti/backoff"
	"github.com/elastic/go-elasticsearch/v7"
	"sync"
	"time"
)

var runOnce sync.Once
var client *elasticsearch.Client

func GetClient() *elasticsearch.Client {
	conf := config.GetElastic()

	runOnce.Do(func() {
		retryBackoff := backoff.NewExponentialBackOff()

		cfg := elasticsearch.Config{
			Addresses:         conf.Addresses,
			Username:          conf.Username,
			Password:          conf.Password,
			RetryOnStatus:     []int{502, 503, 504, 429},
			Logger:            &customLog{},
			EnableDebugLogger: false,
			MaxRetries:        5,
			RetryBackoff: func(i int) time.Duration {
				if i == 1 {
					retryBackoff.Reset()
				}
				return retryBackoff.NextBackOff()
			},
		}

		var err error
		client, err = elasticsearch.NewClient(cfg)
		if err != nil {
			logx.Fatal(err)
			return
		}

		info, err := client.Info()
		if err != nil {
			logx.Fatal(err)
			return
		}

		logx.Infof("elasticsearch is ready status code: %v", info.StatusCode)
	})

	return client
}
