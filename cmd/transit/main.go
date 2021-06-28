package main

import (
	"dora/app/transit"
	"dora/config"
	"flag"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora cmd transit
// 数据接收服务
func main() {
	flag.Parse()
	config.MustLoad(*configFile)

	transit.Serve()
}
