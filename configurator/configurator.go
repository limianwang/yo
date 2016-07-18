package configurator

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Accessor struct {
		AppID     string `json:"app_id"`
		Secret    string `json:"secret"`
		Frequency int    `json:"minutes_to_refresh"`
	} `json:"accessor"`

	Port string `json:"port"`
}

func Load(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	if err := json.Unmarshal(content, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
