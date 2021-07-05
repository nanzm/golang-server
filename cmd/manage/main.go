package main

import (
	"dora/app/manage"
	"dora/config"
	"flag"
	"fmt"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora manage
// 后台管理服务
func main() {
	fmt.Print(`
========================================
++++++++++++++++++++++++++++++++++++++++
           manage server
++++++++++++++++++++++++++++++++++++++++
========================================

`)
	flag.Parse()
	config.MustLoad(*configFile)

	manage.Serve()
}
