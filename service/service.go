package service

import (
	"fmt"

	"github.com/kataras/iris"

	"github.com/limianwang/yo/configurator"
	"github.com/limianwang/yo/service/accessor"
)

// InitAndStart initializes and starts the server
func InitAndStart(conf *configurator.Config) {
	fmt.Println("starting...")

	a := accessor.NewAccessWorker(conf.Accessor.AppID, conf.Accessor.Secret, conf.Accessor.Frequency)

	startServer(conf, a)
}

func startServer(conf *configurator.Config, a *accessor.Access) {
	group := iris.Party("/api")
	{
		group.Get("/validate", func(c *iris.Context) {
			c.JSON(iris.StatusOK, iris.Map{
				"status": "ok",
			})
		})
	}

	iris.Listen(conf.Port)
}
