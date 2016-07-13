package service

import (
	"fmt"

	"github.com/limianwang/yo/configurator"
	"github.com/limianwang/yo/service/accessor"
)

func InitAndStart(conf *configurator.Config) {
	fmt.Println("starting...")
	a := accessor.NewAccessWorker(conf.AppID, conf.Secret)

	fmt.Println(a)
}
