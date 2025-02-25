package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	configFilePath = "etc/config/config.yml"
)

func LoadConfig() (*Config, error) {
	file := configFilePath
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := Config{}
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
