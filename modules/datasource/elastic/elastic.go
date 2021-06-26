package elastic

import (
	"dora/config"
	"dora/pkg/utils/logx"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"os"
	"sync"
)

var runOnce sync.Once
var client *elasticsearch.Client

func GetClient() *elasticsearch.Client {
	conf := config.GetElastic()

	runOnce.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: conf.Addresses,
			Username:  conf.Username,
			Password:  conf.Password,
			Logger: &estransport.ColorLogger{
				Output: os.Stdout,
			},
			EnableDebugLogger: true,
		}

		var err error
		client, err = elasticsearch.NewClient(cfg)
		if err != nil {
			logx.Fatal(err)
			return
		}

		_, err = client.Info()
		if err != nil {
			logx.Fatal(err)
		}

		logx.Infof("elasticsearch is ready!")
	})

	return client
}
