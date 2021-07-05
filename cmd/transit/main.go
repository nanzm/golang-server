package main

import (
	"dora/app/transit"
	"dora/config"
	"flag"
	"fmt"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora cmd transit
// 数据接收服务
func main() {
	fmt.Print(`
========================================
++++++++++++++++++++++++++++++++++++++++
           transit server
++++++++++++++++++++++++++++++++++++++++
========================================

`)
	flag.Parse()
	config.MustLoad(*configFile)

	transit.Serve()
}
