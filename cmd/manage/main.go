package main

import (
	"dora/internal/apps/manage"
	"dora/internal/config"
	"flag"
)

var configFile = flag.String("f", "./config.yml", "the config file")

// dora cmd manage
func main() {
	flag.Parse()
	config.MustLoad(*configFile)

	manage.Serve()
}
