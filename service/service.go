package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/limianwang/yo/config"
	"github.com/limianwang/yo/worker"
)

// InitAndStart initializes and starts the server
func InitAndStart(conf *config.Config) {
	fmt.Println("starting...")

	// a := accessor.NewAccessWorker(conf.Accessor.AppID, conf.Accessor.Secret, conf.Accessor.Frequency)

	startServer(conf)
}

func startServer(conf *config.Config) {
	router := routing.New()

	router.Use(
		access.Logger(log.Printf),
	)

	router.Get("/", func(c *routing.Context) error {
		return c.Write("Root page")
	})

	api := router.Group("/api")
	api.Use(
		content.TypeNegotiator(content.JSON),
	)
	api.Get("/token", func(c *routing.Context) error {
		return c.Write(map[string]string{
			"status": "ok",
			"token":  worker.GetToken(),
		})
	})

	http.Handle("/", router)
	http.ListenAndServe(":10001", nil)
}
