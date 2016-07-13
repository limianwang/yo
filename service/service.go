package service

import (
	"fmt"

	"github.com/limianwang/yo/configurator"
)

func InitAndStart(conf *configurator.Config) {
	fmt.Println("starting...")
	fmt.Println(conf)
}
