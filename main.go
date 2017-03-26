package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/limianwang/yo/config"
	"github.com/limianwang/yo/service"
	"github.com/limianwang/yo/worker"
)

// WebServer that listens for events
// Worker to fetch the accessToken every so often

var (
	configFile = flag.String("config", "config/default.json", "-config /path/to/config.json")
)

func main() {

	flag.Parse()
	fmt.Println(*configFile)

	cfg, err := config.Load(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	go worker.Start("localhost:6379", "")

	fmt.Println("something started...")

	service.InitAndStart(cfg)
}
