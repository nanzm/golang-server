package main

import (
	"dora/app/transit"
	"dora/config"
	"flag"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora cmd transit
func main() {
	flag.Parse()
	config.MustLoad(*configFile)

	transit.Serve()
}
