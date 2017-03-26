package service

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/limianwang/yo/config"
)

// InitAndStart initializes and starts the server
func InitAndStart(conf *config.Config) {
	fmt.Println("starting...")

	// a := accessor.NewAccessWorker(conf.Accessor.AppID, conf.Accessor.Secret, conf.Accessor.Frequency)

	startServer(conf)
}

func startServer(conf *config.Config) {
	group := iris.Party("/api")
	{
		group.Get("/validate", func(c *iris.Context) {
			c.JSON(iris.StatusOK, iris.Map{
				"status": "ok",
			})
		})
	}

	iris.Listen(fmt.Sprintf(":%s", conf.Port))
}
