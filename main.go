package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

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

	cfg, err := config.Load(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	w := worker.Init(
		cfg.Redis.Host,
		cfg.Redis.Password,
		cfg.Redis.DB,
		cfg.Accessor.AppID,
		cfg.Accessor.Secret,
		cfg.Accessor.Frequency,
	)

	done := make(chan bool)

	go w.Start(done)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Print("Caught SIGINT.... passing end")
		done <- true

		select {
		case <-done:
			os.Exit(0)
		}
	}()

	log.Println("Starting server....")

	service.InitAndStart(cfg)
}
