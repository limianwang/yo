package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
)

type Config struct{}

func main() {
	app := cli.NewApp()
	app.Name = "yo"

	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		loadConfig(c.String("config"))
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load Configuration from `FILE`",
			Value: "config/default.json",
		},
	}

	app.Run(os.Args)
}

func loadConfig(s string) (*Config, error) {
	content, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := json.Unmarshal(content, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
