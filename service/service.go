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

	a := accessor.NewAccessWorker(conf.AppID, conf.Secret)

	startServer(conf)

	fmt.Println(a)
}

func startServer(conf *configurator.Config) {
	fmt.Println("Hello...")

	iris.Listen(conf.Port)
}
