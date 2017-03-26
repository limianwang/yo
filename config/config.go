package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config object that hosts the json formatted variables
type Config struct {
	Accessor struct {
		AppID     string `json:"app_id"`
		Secret    string `json:"secret"`
		Frequency int    `json:"minutes_to_refresh"`
	} `json:"accessor"`

	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis"`
}

// Load initializes the config object based on a file path
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
