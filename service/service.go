package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/limianwang/yo/config"
)

type Tokener interface {
	GetToken() (string, error)
}

// InitAndStart initializes and starts the server
func InitAndStart(conf *config.Config, w Tokener) {
	fmt.Println("starting...")

	// a := accessor.NewAccessWorker(conf.Accessor.AppID, conf.Accessor.Secret, conf.Accessor.Frequency)

	startServer(conf, w)
}

func startServer(conf *config.Config, w Tokener) {
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
		token, err := w.GetToken()

		if err != nil {
			return c.Write(map[string]string{
				"status": "not_ok",
			})
		}

		return c.Write(map[string]string{
			"status": "ok",
			"token":  token,
		})
	})

	http.Handle("/", router)
	http.ListenAndServe(":10001", nil)
}
