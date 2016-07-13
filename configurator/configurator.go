package configurator

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	AppID  string `json:"app_id"`
	Secret string `json:"secret"`
}

func LoadConfig(path string) (*Config, error) {
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
