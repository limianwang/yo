package main

import (
	"os"

	"github.com/limianwang/yo/configurator"
	"github.com/limianwang/yo/service"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "yo"

	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		config, _ := configurator.Load(c.String("config"))
		service.InitAndStart(config)
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load Configuration from `FILE`",
			Value: "configurator/default.json",
		},
	}

	app.Run(os.Args)
}
