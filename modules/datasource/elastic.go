package datasource

import (
	"dora/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"os"
	"sync"
)

var runOnce sync.Once
var client *elasticsearch.Client

func GetElasticClient() *elasticsearch.Client {
	conf := config.GetElastic()

	runOnce.Do(func() {
		cfg := elasticsearch.Config{
			Addresses:         conf.Addresses,
			Username:          conf.Username,
			Password:          conf.Password,
			Logger:            &estransport.ColorLogger{Output: os.Stdout},
			EnableDebugLogger: true,
		}

		var err error
		client, err = elasticsearch.NewClient(cfg)
		if err != nil {
			//logx.Error(err)
		}
		//res, err := client.Info()
		//if err != nil {
		//	panic(err)
		//}
		//logx.Printf("%v \n", res)
	})

	return client
}

func CloseElasticClient() {
}
