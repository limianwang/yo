package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/limianwang/yo/configurator"
)

func main() {
	app := cli.NewApp()
	app.Name = "yo"

	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		config, _ := configurator.LoadConfig(c.String("config"))
		fmt.Println(config)
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
