package config

import "github.com/huyrun/go-admin/modules/config"

type Config struct {
	ServerAddress  string `yaml:"server_address"`
	*config.Config `yaml:",inline"`
}
