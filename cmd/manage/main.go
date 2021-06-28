package main

import (
	"dora/app/manage"
	"dora/config"
	"flag"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora manage
// 后台管理服务
func main() {
	flag.Parse()
	config.MustLoad(*configFile)

	manage.Serve()
}
