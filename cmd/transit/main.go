package main

import (
	"dora/internal/apps/transit"
	"dora/internal/config"
	"flag"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora cmd transit
func main() {
	flag.Parse()
	config.MustLoad(*configFile)

	transit.Serve()
}
